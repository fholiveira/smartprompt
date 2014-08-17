package parsers

import (
	"regexp"
	"strings"
)

type PromptLine struct {
	Text string
}

func (prompt *PromptLine) Apply(token Token, value string) {
	tokenKey := "{" + string(token) + "}"
	prompt.Text = strings.Replace(prompt.Text, tokenKey, value, -1)
}

func (prompt PromptLine) Tokens() []Token {
	regex := regexp.MustCompile("{([^}]+)}")
	matches := regex.FindAllStringSubmatch(prompt.Text, -1)

	column := make([]Token, 0)
	for _, row := range matches {
		column = append(column, Token(row[1]))
	}

	return column
}
