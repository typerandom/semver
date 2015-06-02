package semver_test

import (
	"github.com/typerandom/semver"
	"testing"
)

func assertThatInvalidVersionIsInvalid(t *testing.T, value string) {
	if _, err := semver.Parse(value); err == nil {
		t.Error("Expected error, but didn't get any.")
	}
}

func assertThatVersionIsValid(t *testing.T, value string, expectedMajor int, expectedMinor int, expectedPatch int, expectedPreRelease []string, expectedBuild []string) {
	v, err := semver.Parse(value)

	if err != nil {
		t.Errorf("Didn't expect error, but got '%s'.", err)
	}

	if v.Major() != expectedMajor {
		t.Errorf("Expected major version to be %d, but it was %d.", expectedMajor, v.Major())
	}

	if v.Minor() != expectedMinor {
		t.Errorf("Expected minor version to be %d, but it was %d.", expectedMinor, v.Minor())
	}

	if v.Patch() != expectedPatch {
		t.Errorf("Expected patch version to be %d, but it was %d.", expectedPatch, v.Patch())
	}

	if len(v.PreRelease()) != len(expectedPreRelease) {
		t.Errorf("Expected build metadata length to be %d, but it was %d.", len(expectedPreRelease), len(v.PreRelease()))
	}

	for i, metadata := range expectedPreRelease {
		if metadata != v.PreRelease()[i] {
			t.Errorf("Expected pre-release metadata[%d] to be '%s', but it was '%s'.", i, metadata, v.PreRelease()[i])
		}
	}

	if len(v.Build()) != len(expectedBuild) {
		t.Errorf("Expected build metadata length to be %d, but it was %d.", len(expectedBuild), len(v.Build()))
	}

	for i, metadata := range expectedBuild {
		if metadata != v.Build()[i] {
			t.Errorf("Expected build metadata[%d] to be '%s', but it was '%s'.", i, metadata, v.Build()[i])
		}
	}

	if value != v.String() {
		t.Errorf("Expected string version to be '%s', but it was '%s'.", value, v.String())
	}
}

func TestThatValidVersionIsValid(t *testing.T) {
	assertThatVersionIsValid(t, "1.2.3", 1, 2, 3, []string{}, []string{})
}

func TestThatValidVersionWithPreReleaseMetadataIsValid(t *testing.T) {
	assertThatVersionIsValid(t, "1.2.3-early-bird.135", 1, 2, 3, []string{"early-bird", "135"}, []string{})
}

func TestThatValidVersionWithBuildMetadataIsValid(t *testing.T) {
	assertThatVersionIsValid(t, "1.2.3+early-bird.135", 1, 2, 3, []string{}, []string{"early-bird", "135"})
}

func TestThatValidVersionWithPreReleaseAndBuildMetadataIsValid(t *testing.T) {
	assertThatVersionIsValid(t, "1.2.3-super-bird.951+early-bird.135", 1, 2, 3, []string{"super-bird", "951"}, []string{"early-bird", "135"})
}

func TestThatVariousValidVersionsAreValid(t *testing.T) {
	assertThatVersionIsValid(t, "1.0.0", 1, 0, 0, []string{}, []string{})
	assertThatVersionIsValid(t, "0.1.0", 0, 1, 0, []string{}, []string{})
	assertThatVersionIsValid(t, "0.0.1", 0, 0, 1, []string{}, []string{})
	assertThatVersionIsValid(t, "1.0.0-beta+exp.sha.5114f85", 1, 0, 0, []string{"beta"}, []string{"exp", "sha", "5114f85"})
}
func TestThatVariousInvalidVersionsAreInvalid(t *testing.T) {
	assertThatInvalidVersionIsInvalid(t, "")
	assertThatInvalidVersionIsInvalid(t, "0.0.0")
	assertThatInvalidVersionIsInvalid(t, "0.0.-1")
	assertThatInvalidVersionIsInvalid(t, "0.-1.0")
	assertThatInvalidVersionIsInvalid(t, "-1.0.0")
	assertThatInvalidVersionIsInvalid(t, "1.0.0.0")
	assertThatInvalidVersionIsInvalid(t, "...")
}
