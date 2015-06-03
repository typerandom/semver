package semver_test

import (
	"github.com/typerandom/semver"
	"testing"
)

func TestThatWhenVersionsAreSortedTheyAreReturnedInRightOrder(t *testing.T) {
	versions := []semver.Version{
		semver.New("0.1.5-beta"),
		semver.New("1.0.0-alpha"),
		semver.New("0.1.0"),
		semver.New("1.0.0"),
		semver.New("0.1.5-alpha.1"),
		semver.New("0.1.0-alpha"),
		semver.New("1.0.0-alpha.1"),
		semver.New("1.0.0-beta"),
		semver.New("0.1.5-alpha"),
	}

	semver.Sort(versions)

	expectedVersionOrder := []string{
		"0.1.0-alpha",
		"0.1.0",
		"0.1.5-alpha",
		"0.1.5-alpha.1",
		"0.1.5-beta",
		"1.0.0-alpha",
		"1.0.0-alpha.1",
		"1.0.0-beta",
		"1.0.0",
	}

	if len(expectedVersionOrder) != len(versions) {
		t.Errorf("Expected sorted array to have %d versions, but there were %d.", len(expectedVersionOrder), len(versions))
	}

	for i, version := range expectedVersionOrder {
		if version != versions[i].String() {
			t.Errorf("Expected version at offset %d to be '%s', but it was '%s'.", i, version, versions[i].String())
		}
	}
}
