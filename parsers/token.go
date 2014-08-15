package parsers

import "strings"

type Token string

func (token Token) Name() string {
	return strings.Split(string(token), "|")[0]
}

func (token Token) Parameter() string {
	values := strings.Split(string(token), "|")

	if len(values) > 1 {
		return values[1]
	}

	return ""
}
