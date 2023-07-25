package tools

import "regexp"

func FindByRegx(source string, regexstr string) []string {
	regex := regexp.MustCompile(regexstr)
	matches := regex.FindAllString(source, -1)
	return matches
}
