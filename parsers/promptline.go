package parsers

import (
	"regexp"
	"strings"
)

type PromptLine struct {
	Text string
}

func (prompt *PromptLine) Apply(token Token, value string) {
	tokenKey := "{" + token.text + "}"
	prompt.Text = strings.Replace(prompt.Text, tokenKey, value, -1)
}

func (prompt PromptLine) Tokens(separator string) []Token {
	regex := regexp.MustCompile("{([^}]+)}")
	matches := regex.FindAllStringSubmatch(prompt.Text, -1)

	column := make([]Token, 0)
	for _, row := range matches {
		column = append(column, NewToken(row[1], separator))
	}

	return column
}
