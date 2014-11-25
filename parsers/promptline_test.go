package parsers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPromptLineTokens(t *testing.T) {
	prompt := PromptLine{"{RED} {BLUE:bold}/{GREEN}"}
	expectedTokens := []Token{
		NewToken("RED", ":"),
		NewToken("BLUE:bold", ":"),
		NewToken("GREEN", ":")}

	assert.Equal(t, prompt.Tokens(":"), expectedTokens)
}

func TestGetPromptLineTokensWhenSomeTokensHaveParameters(t *testing.T) {
	prompt := PromptLine{"{RED:bold} {time|mm/yy}/{GREEN}"}
	expectedTokens := []Token{
		NewToken("RED:bold", ":"),
		NewToken("time|mm/yy", ":"),
		NewToken("GREEN", ":")}

	assert.Equal(t, prompt.Tokens(":"), expectedTokens)
}

func TestApplyTokenInPromptLine(t *testing.T) {
	prompt := PromptLine{"{RED} {user}@{host}"}
	token := NewToken("user", "|")

	prompt.Apply(token, "username123")

	assert.Equal(t, "{RED} username123@{host}", prompt.Text)
}

func TestApplyTokeniWithParameterInPromptLine(t *testing.T) {
	prompt := PromptLine{"{time|yy/mm} {user}@{host}"}
	token := NewToken("time|yy/mm", "|")

	prompt.Apply(token, "05/1992")

	assert.Equal(t, "05/1992 {user}@{host}", prompt.Text)
}
