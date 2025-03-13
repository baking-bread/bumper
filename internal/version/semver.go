package version

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type SemVer struct {
	Major      int
	Minor      int
	Patch      int
	PreRelease string // optional
	Build      string // optional
}

// https://semver.org/#is-there-a-suggested-regular-expression-regex-to-check-a-semver-string
// https://regex101.com/
var semverRegexp = regexp.MustCompile(`(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?`)

func NewSemVer(major int, minor int, patch int, pre_release string, build string) *SemVer {
	return &SemVer{
		Major:      major,
		Minor:      minor,
		Patch:      patch,
		PreRelease: pre_release,
		Build:      build,
	}
}

func Parse(version string) (*SemVer, error) {
	var result = &SemVer{}

	// parse using regexp

	// matches only the left most version matched
	// should we match all? and check if we only found one version?
	matches := semverRegexp.FindStringSubmatch(version)
	if matches == nil {
		return nil, fmt.Errorf("could not parse version from %s", version)
	}

	// regex should make sure, that there is no error
	var err error
	if result.Major, err = strconv.Atoi(matches[1]); err != nil {
		return nil, fmt.Errorf("could not parse major version from %s", version)
	}
	if result.Minor, err = strconv.Atoi(matches[2]); err != nil {
		return nil, fmt.Errorf("could not parse pathc version from %s", version)
	}
	if result.Patch, err = strconv.Atoi(matches[3]); err != nil {
		return nil, fmt.Errorf("could not parse patch version from %s", version)
	}
	result.PreRelease = matches[4]
	result.Build = matches[5]

	return result, nil
}

func Replace(version string, update string) string {
	return semverRegexp.ReplaceAllString(version, update)
}

// compares version by major, minor and patch numbers
// does not compare with pre releases
func (sv *SemVer) Compare(other *SemVer) int {
	if sv.Major != other.Major {
		return sv.Major - other.Major
	}
	if sv.Minor != other.Minor {
		return sv.Minor - other.Minor
	}
	return sv.Patch - other.Patch
}

func (sv *SemVer) Bump(bumpType string) error {
	switch bumpType {
	case "patch":
		sv.Patch++
	case "minor":
		sv.Minor++
		sv.Patch = 0
	case "major":
		sv.Major++
		sv.Minor = 0
		sv.Patch = 0
	default:
		return errors.New("invalid bump type")
	}

	return nil
}

func (sv *SemVer) String() string {
	return fmt.Sprintf("%d.%d.%d%s%s", sv.Major, sv.Minor, sv.Patch, sv.PreRelease, sv.Build)
}
