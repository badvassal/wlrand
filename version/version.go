package version

import "fmt"

const Version = "0.0.2"

var (
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
