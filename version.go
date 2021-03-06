package sequence

import (
	"fmt"
	"runtime/debug"
)

var (
	Version = "0.0.0"

	Prerelease = ""
)

func Semver() string {
	version := Version

	if Prerelease != "" {
		version = fmt.Sprintf("%s-%s", version, Prerelease)
	}

	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		var (
			revision string
			modified bool
		)
		for _, setting := range buildInfo.Settings {
			switch setting.Key {
			case "vcs.revision":
				revision = setting.Value
			case "vcs.modified":
				modified = setting.Value == "true"
			}
		}

		if revision != "" {
			i := len(revision)
			if i > 7 {
				i = 7
			}
			version = fmt.Sprintf("%s+%s", version, revision[:i])
		}

		if modified {
			version = fmt.Sprintf("%s*", version)
		}
	}

	return version
}
