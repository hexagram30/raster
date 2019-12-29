package common

import "github.com/geomyidia/util/version"

// Versioning data
var (
	vsn        string
	buildDate  string
	gitCommit  string
	gitBranch  string
	gitSummary string
)

// VersionData stuff for things
func VersionData() *version.ProjectVersion {
	return &version.ProjectVersion{
		Semantic:   vsn,
		BuildDate:  buildDate,
		GitCommit:  gitCommit,
		GitBranch:  gitBranch,
		GitSummary: gitSummary,
	}
}

// BuildString ...
func BuildString() string {
	return version.BuildString(VersionData())
}

// VersionString ...
func VersionString() string {
	return version.String(VersionData())
}

// VersionedBuildString ...
func VersionedBuildString() string {
	return version.VersionedBuildString(VersionData())
}
