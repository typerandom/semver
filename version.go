package semver

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Represents a semantic version.
// http://semver.org/

// Regex used for parsing a semantic version (2.0) as specified by http://semver.org/
var version20Regexp = regexp.MustCompile("^(\\d+)\\.(\\d+)\\.(\\d+)(\\-(([0-9A-Za-z-]+)(\\.)?)+)?(\\+(([0-9A-Za-z-]+)(\\.)?)+)?$")

// Version represents the structure of the Semantic Versioning 2.0 scheme.
type Version struct {
	major      int
	minor      int
	patch      int
	preRelease []string
	build      []string
}

// Compare two semantic versions. This can be used for sort methods.
func Compare(a Version, b Version) int {
	if a.Before(b) {
		return -1
	} else if a.After(b) {
		return 1
	} else {
		return 0
	}
}

// Parse metadata such as pre-release and build of a version.
func parseMetadata(metadata string) ([]string, error) {
	if len(metadata) == 0 {
		return []string{}, nil
	}

	if metadata[0] != '-' && metadata[0] != '+' {
		return nil, errors.New("Invalid metadata indicator sign '" + string(metadata[0]) + "'.")
	}

	if metadata[len(metadata)-1] == '.' {
		return nil, errors.New("Metadata cannot end with dot.")
	}

	return strings.Split(metadata[1:], "."), nil
}

// Parse tries to parse a raw value. Returns error if it fails.
func Parse(value string) (*Version, error) {
	groups := version20Regexp.FindAllStringSubmatch(value, -1)

	if len(groups) == 0 {
		return nil, errors.New("Invalid version format.")
	}

	matches := groups[0]

	major, _ := strconv.ParseInt(matches[1], 10, 32)
	minor, _ := strconv.ParseInt(matches[2], 10, 32)
	patch, _ := strconv.ParseInt(matches[3], 10, 32)

	preRelease, err := parseMetadata(matches[4])

	if err != nil {
		return nil, errors.New("Invalid version format.")
	}

	build, err := parseMetadata(matches[8])

	if err != nil {
		return nil, errors.New("Invalid version format.")
	}

	// Version cannot be zero.
	if (major + minor + patch) == 0 {
		return nil, errors.New("Invalid version format.")
	}

	return &Version{
		major:      int(major),
		minor:      int(minor),
		patch:      int(patch),
		preRelease: preRelease,
		build:      build,
	}, nil
}

// New creates a new version given a raw value. Panics if wrong format.
func New(version string) Version {
	result, err := Parse(version)

	if err != nil {
		panic(err.Error())
	}

	return *result
}

// Major gets the major version.
func (v Version) Major() int {
	return v.major
}

// Minor gets the minor version.
func (v Version) Minor() int {
	return v.minor
}

// Patch gets the patch version.
func (v Version) Patch() int {
	return v.patch
}

// PreRelease gets the pre-release build metadata.
func (v Version) PreRelease() []string {
	return v.preRelease
}

// Build gets the build metadata.
func (v Version) Build() []string {
	return v.build
}

// Same determines whether or not this version is equal to another version.
func (v Version) Same(c Version) bool {
	return !v.Before(c) && !v.After(c)
}

// Before determines whether or this version is a precursor to another version.
func (v Version) Before(t Version) bool {
	if v.major < t.major {
		return true
	} else if v.minor < t.minor {
		return true
	} else if v.patch < t.patch {
		return true
	}
	return false
}

// After determines whether or this version is a successor to another version.
func (v Version) After(t Version) bool {
	if v.major > t.major {
		return true
	} else if v.minor > t.minor {
		return true
	} else if v.patch > t.patch {
		return true
	}
	return false
}

// String gets the string representation of this version.
func (v *Version) String() string {
	result := fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.patch)

	if len(v.preRelease) > 0 {
		result += "-" + strings.Join(v.preRelease, ".")
	}

	if len(v.build) > 0 {
		result += "+" + strings.Join(v.build, ".")
	}

	return result
}
