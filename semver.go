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

// Version represents a Semantic Version.
type Version interface {
	// Major gets the major version.
	Major() int

	// Minor gets the minor version.
	Minor() int

	// Patch gets the patch version.
	Patch() int

	// PreRelease gets the pre-release metadata.
	PreRelease() []string

	// Build gets the build metadata.
	Build() []string

	// Same determines whether or not this version is equal to another version. Note: build metadata may differ.
	Same(v Version) bool

	// Before determines whether or not this version is a precursor to another version.
	Before(v Version) bool

	// After determines whether or not this version is a successor to another version.
	After(v Version) bool

	// String gets the string representation of this version.
	String() string
}

// Regex used for parsing a semantic version (2.0) as specified by http://semver.org/
var version20Regexp = regexp.MustCompile("^(\\d+)\\.(\\d+)\\.(\\d+)(\\-(([0-9A-Za-z-]+)(\\.)?)+)?(\\+(([0-9A-Za-z-]+)(\\.)?)+)?$")

// Version represents the structure of the Semantic Versioning 2.0 scheme.
type version20 struct {
	major      int
	minor      int
	patch      int
	preRelease []string
	build      []string
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
func Parse(value string) (Version, error) {
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

	return &version20{
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

	return result
}

func (v *version20) Major() int {
	return v.major
}

func (v *version20) Minor() int {
	return v.minor
}

func (v *version20) Patch() int {
	return v.patch
}

func (v *version20) PreRelease() []string {
	return v.preRelease
}

func (v *version20) Build() []string {
	return v.build
}

// Compares pre-releases from one version with pre-releases of another.
func comparePreReleases(a []string, b []string) int {
	lenA := len(a)
	lenB := len(b)

	if lenA == 0 && lenB == 0 {
		return 0
	} else if lenA == 0 {
		return 1
	} else if lenB == 0 {
		return -1
	}

	lim := lenA

	if lenB < lenA {
		lim = lenB
	}

	for i := 0; i < lim; i++ {
		preA := a[i]
		preB := b[i]
		if preA == preB {
			continue
		} else if preA > preB {
			return 1
		} else { // preA < preB
			return -1
		}
	}

	if lenA > lenB {
		return 1
	}

	return 0
}

// Compares two versions and returns an int indicating the relation of A to B.
// The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
func compareVersions(a Version, b Version) int {
	if a.Major() < b.Major() {
		return -1
	} else if a.Major() > b.Major() {
		return 1
	}

	if a.Minor() < b.Minor() {
		return -1
	} else if a.Minor() > b.Minor() {
		return 1
	}

	if a.Patch() < b.Patch() {
		return -1
	} else if a.Patch() > b.Patch() {
		return 1
	}

	return comparePreReleases(a.PreRelease(), b.PreRelease())
}

func (v *version20) Same(t Version) bool {
	return compareVersions(v, t) == 0
}

func (v *version20) Before(t Version) bool {
	return compareVersions(v, t) < 0
}

func (v *version20) After(t Version) bool {
	return compareVersions(v, t) > 0
}

func (v *version20) String() string {
	result := fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.patch)

	if len(v.preRelease) > 0 {
		result += "-" + strings.Join(v.preRelease, ".")
	}

	if len(v.build) > 0 {
		result += "+" + strings.Join(v.build, ".")
	}

	return result
}
