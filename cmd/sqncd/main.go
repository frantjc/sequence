package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"

	"github.com/frantjc/sequence"
	"github.com/frantjc/sequence/conf"
	"github.com/frantjc/sequence/log"
	"github.com/frantjc/sequence/meta"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:               fmt.Sprintf("%sd", meta.Name),
	Version:           meta.Semver(),
	PersistentPreRunE: persistentPreRun,
	RunE:              run,
}

var (
	configFilePath string
	verbose        bool
	socket         string
	port           int
	rootDir        string
	stateDir       string
	workDir        string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&configFilePath, "config", "", "config file")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "verbose")
	rootCmd.PersistentFlags().StringVar(&socket, "sock", "", "unix socket")
	rootCmd.PersistentFlags().IntVar(&port, "port", 0, "port")
	rootCmd.PersistentFlags().StringVar(&rootDir, "root-dir", "", "root dir")
	rootCmd.PersistentFlags().StringVar(&stateDir, "state-dir", "", "state dir")
	wd, _ := os.Getwd()
	rootCmd.PersistentFlags().StringVar(&workDir, "context", wd, "context")
}

func newConfig() (*conf.Config, error) {
	configOpts := []conf.ConfigOpt{
		conf.WithConfig(&conf.Config{
			Verbose:  verbose,
			Socket:   socket,
			Port:     port,
			RootDir:  rootDir,
			StateDir: stateDir,
		}),
	}

	if configFilePath != "" {
		configOpts = append(configOpts, conf.WithConfigFilePath(workDir, configFilePath))
	}

	configOpts = append(configOpts, conf.WithConfigFromEnv)

	if _, err := os.Stat(conf.DefaultUserConfigFilePath); err == nil {
		configOpts = append(configOpts, conf.WithConfigFilePath(workDir, conf.DefaultUserConfigFilePath))
	}

	configOpts = append(configOpts, conf.WithDefaultUserConfig)

	if _, err := os.Stat(conf.DefaultSystemConfigFilePath); err == nil {
		configOpts = append(configOpts, conf.WithConfigFilePath(workDir, conf.DefaultSystemConfigFilePath))
	}

	configOpts = append(configOpts, conf.WithDefaultSystemConfig)

	return conf.New(configOpts...)
}

func init() {
	rootCmd.SetVersionTemplate(
		fmt.Sprintf("{{ with .Name }}{{ . }}{{ end }}{{ with .Version }}{{ . }}{{ end }} %s\n", runtime.Version()),
	)
}

func persistentPreRun(cmd *cobra.Command, args []string) error {
	c, err := newConfig()
	if err != nil {
		return err
	}
	log.SetVerbose(c.Verbose)
	return nil
}

func run(cmd *cobra.Command, args []string) error {
	var (
		ctx    = cmd.Context()
		c, err = newConfig()
	)
	if err != nil {
		return err
	}

	addr := strings.TrimPrefix(c.Address(), "unix://")
	os.Remove(addr)
	defer os.Remove(addr)

	l, err := net.Listen(c.Network(), addr)
	if err != nil {
		return err
	}

	s, err := sequence.NewServer(ctx, sequence.WithAnyRuntime)
	if err != nil {
		return err
	}

	return s.Serve(l)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
