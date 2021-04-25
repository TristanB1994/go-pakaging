package gostrings

import "regexp"

func Squish(text string) string {
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(text, " ")
}
