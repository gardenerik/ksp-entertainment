package music

import (
	"regexp"
	"strings"
)

func replaceRegex(regex, in, replacement string) string {
	return regexp.MustCompile(regex).ReplaceAllString(in, replacement)
}

func ParseTitle(title string) (name string, artist string) {
	title = replaceRegex("(?i)\\(.*video.*\\)", title, "")
	title = replaceRegex("(?i)\\(.*live.*\\)", title, "")
	title = replaceRegex("(?i)\\(.*official.*\\)", title, "")

	title = replaceRegex("\\s\\s", title, " ")

	// best effort, most common format Artist - Name
	r := regexp.MustCompile("[\\|-]")
	if r.MatchString(title) {
		parts := r.Split(title, -1)
		if len(parts) == 2 {
			return strings.TrimSpace(parts[1]), strings.TrimSpace(parts[0])
		}
	}
	return title, ""
}
