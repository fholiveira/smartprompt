package parsers

import "regexp"

type Parser interface {
	Parse(prompt string) (string, error)
}

func getTokens(prompt string) []string {
	regex := regexp.MustCompile("{([^}]+)}")
	matches := regex.FindAllStringSubmatch(prompt, -1)

	column := make([]string, 0)
	for _, row := range matches {
		column = append(column, row[1])
	}

	return column
}
