package version

import "fmt"

var (
	Version    string = "unknown_version"
	BuildDate  string = "unknown_date"
	CommitHash string = "unknown_commit"
	GitState   string = ""
)

func VersionStr() string {
	gitStr := CommitHash
	if GitState != "" {
		gitStr += "-" + GitState
	}
	return fmt.Sprintf("%s / %s / %s", Version, gitStr, BuildDate)
}
