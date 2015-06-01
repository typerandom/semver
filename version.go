package version

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Represents a semantic Semantic
// http://semver.org/

type Semantic struct {
	major int
	minor int
	patch int
}

var namedParts = []string{"major", "minor", "patch"}

func Parse(version string) (*Semantic, error) {
	parts := strings.SplitN(version, ".", 3)

	if len(parts) != 3 {
		return nil, errors.New("A semantic version must only consist of 3 parts. Major, minor and patch.")
	}

	var parsed [3]int

	for i, part := range parts {
		parsedPart, err := strconv.ParseInt(part, 10, 32)

		if err != nil {
			return nil, errors.New("Semantic version part '" + namedParts[i] + "' must be a valid integer.")
		}

		if parsedPart < 0 {
			return nil, errors.New("Semantic version part '" + namedParts[i] + "' cannot be negative.")
		}

		parsed[i] = int(parsedPart)
	}

	return &Semantic{
		major: parsed[0],
		minor: parsed[1],
		patch: parsed[2],
	}, nil
}

func New(version string) Semantic {
	result, err := Parse(version)

	if err != nil {
		panic(err.Error())
	}

	return *result
}

func (v Semantic) Valid() bool {
	return v.major > 0 || v.minor > 0 || v.patch > 0
}

func (v Semantic) Major() int {
	return v.major
}

func (v Semantic) IncrementMajor() Semantic {
	nv := v.Copy()
	nv.major++
	return nv
}

func (v Semantic) Minor() int {
	return v.minor
}

func (v Semantic) IncrementMinor() Semantic {
	nv := v.Copy()
	nv.minor++
	return nv
}

func (v Semantic) Patch() int {
	return v.patch
}

func (v Semantic) IncrementPatch() Semantic {
	nv := v.Copy()
	nv.patch++
	return nv
}

func (v Semantic) IsSame(c Semantic) bool {
	return !v.IsBefore(c) && !v.IsAfter(c)
}

func (v Semantic) IsBefore(t Semantic) bool {
	if v.major < t.major {
		return true
	} else if v.minor < t.minor {
		return true
	} else if v.patch < t.patch {
		return true
	}
	return false
}

func (v Semantic) IsAfter(t Semantic) bool {
	if v.major > t.major {
		return true
	} else if v.minor > t.minor {
		return true
	} else if v.patch > t.patch {
		return true
	}
	return false
}

func (v Semantic) Copy() Semantic {
	return Semantic{
		major: v.major,
		minor: v.minor,
		patch: v.patch,
	}
}

func (v *Semantic) String() string {
	return fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.patch)
}
