package version

import (
	"fmt"
	"regexp"
)

func Match(versionRegexp *regexp.Regexp, content string) (string, error) {
	matches := versionRegexp.FindStringSubmatch(content)

	if matches == nil {
		return "", fmt.Errorf("could not find version in %s", content)
	}

	return matches[1], nil
}

func Update(versionRegexp *regexp.Regexp, content string, version string) string {
	match, _ := Match(versionRegexp, content)
	update := Replace(match, version)

	fmt.Println(match)
	fmt.Println(update)

	versionRegexp.ReplaceAllString(content, `$version`)

	return versionRegexp.ReplaceAllString(content, update)
}
