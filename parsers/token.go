package parsers

import "strings"

type Token struct {
	text      string
	separator string
}

func (token Token) Name() string {
	return strings.Split(token.text, token.separator)[0]
}

func (token Token) Parameters() []string {
	values := strings.Split(token.text, token.separator)

	if len(values) > 1 {
		return values[1:]
	}

	return nil
}

func NewToken(text string, separator string) Token {
	return Token{text, separator}
}
