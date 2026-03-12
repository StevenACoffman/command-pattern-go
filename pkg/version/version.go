package version

import "fmt"

// Populated during build or runtime init.
var (
	AppName      = "unknown" // What the app is called
	Project      = "unknown" // Which GCP Project e.g. khan-academy
	Date         = ""        // Build Date in RFC3339 e.g. $(date -u +"%Y-%m-%dT%H:%M:%SZ")
	GitCommit    = "?"       // Git Commit sha1 of source
	Version      = "v0.0.0"  // go mod goVersion: v0.0.0-20200214070026-92e9ce6ff79f
	HumanVersion = fmt.Sprintf(
		"%s %s %s (%s) on %s",
		AppName,
		Project,
		Version,
		GitCommit,
		Date,
	)
)

func init() {
	info := getBuildInfo()
	if info == nil {
		return
	}
	if info.Main.Version != "" {
		Version = info.Main.Version
	}
	for _, kv := range info.Settings {
		if kv.Value == "" {
			continue
		}
		switch kv.Key {
		case "vcs.revision":
			GitCommit = kv.Value
		case "vcs.time":
			Date = kv.Value
		}
	}
}
