package semver

import (
	"sort"
)

type versions []Version

func (s versions) Len() int {
	return len(s)
}

func (s versions) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s versions) Less(i, j int) bool {
	return s[i].Before(s[j])
}

// Sort a list of versions from lowest version to highest.
func Sort(items []Version) {
	sort.Sort(versions(items))
}
