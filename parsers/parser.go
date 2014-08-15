package parsers

import (
	"regexp"
	"strings"
)

type PromptLine struct {
	Text string
}

func (prompt *PromptLine) Apply(token, value string) {
	prompt.Text = strings.Replace(prompt.Text, "{"+token+"}", value, -1)
}

func (prompt *PromptLine) Tokens() []string {
	regex := regexp.MustCompile("{([^}]+)}")
	matches := regex.FindAllStringSubmatch(prompt.Text, -1)

	column := make([]string, 0)
	for _, row := range matches {
		column = append(column, row[1])
	}

	return column
}

type Parser interface {
	Parse(prompt PromptLine) (PromptLine, error)
}
