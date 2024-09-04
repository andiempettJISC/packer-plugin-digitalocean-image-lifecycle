package version

import "github.com/hashicorp/packer-plugin-sdk/version"

var (

	Version string

	// A pre-release marker for the version. If this is "" (empty string)
	// then it means that it is a final release. Otherwise, this is a pre-release
	// such as "dev" (in development), "beta", "rc1", etc.
	VersionPrerelease = ""

	// The metadata for the version, this is optional information to add around
	// a particular release.
	// This has no impact on the ordering of plugins, and is ignored for non-human eyes.
	VersionMetadata = ""	

	PluginVersion     = version.NewPluginVersion(Version, VersionPrerelease, VersionMetadata)
)
