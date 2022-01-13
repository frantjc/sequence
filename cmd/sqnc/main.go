package main

import (
	"context"
	"fmt"
	"os"
	"runtime"

	"github.com/frantjc/sequence/meta"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:               meta.Name,
	Version:           meta.Semver(),
	PersistentPreRunE: persistentPreRun,
}

var (
	verbose bool
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose")
	rootCmd.AddCommand(
		runCmd,
		pluginCmd,
	)
	rootCmd.SetVersionTemplate(
		fmt.Sprintf("{{ with .Name }}{{ . }}{{ end }}{{ with .Version }}{{ . }}{{ end }} %s\n", runtime.Version()),
	)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}

func persistentPreRun(cmd *cobra.Command, args []string) error {
	if verbose {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	return nil
}
