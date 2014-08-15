package parsers

import "regexp"

type Parser interface {
	Parse(prompt string) (string, error)
}

func getTokens(prompt string) []string {
	regex := regexp.MustCompile("{([^}]+)}")
	return regex.FindAllString(prompt, -1)
}
