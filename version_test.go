package version_test

import (
	"github.com/typerandom/version"
	"testing"
)

func TestThatValidSemanticVersionCanBeCreated(t *testing.T) {
	v := version.New("1.2.3")

	if v.Major() != 1 {
		t.Errorf("Expected major version to be 1, but it was %d.", v.Major())
	}

	if v.Minor() != 2 {
		t.Errorf("Expected minor version to be 1, but it was %d.", v.Minor())
	}

	if v.Patch() != 3 {
		t.Errorf("Expected patch version to be 1, but it was %d.", v.Patch())
	}
}
