package parsers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPromptLineTokens(t *testing.T) {
	prompt := PromptLine{"{RED} {BLUE:bold}/{GREEN}"}
	expectedTokens := []Token{Token("RED"), Token("BLUE:bold"), Token("GREEN")}
	assert.Equal(t, prompt.Tokens(), expectedTokens)
}

func TestGetPromptLineTokensWhenSomeTokensHaveParameters(t *testing.T) {
	prompt := PromptLine{"{RED:bold} {time|mm/yy}/{GREEN}"}
	expectedTokens := []Token{Token("RED:bold"), Token("time|mm/yy"), Token("GREEN")}
	assert.Equal(t, prompt.Tokens(), expectedTokens)
}

func TestApplyTokenInPromptLine(t *testing.T) {
	prompt := PromptLine{"{RED} {user}@{host}"}
	token := Token("user")

	prompt.Apply(token, "username123")

	assert.Equal(t, "{RED} username123@{host}", prompt.Text)
}

func TestApplyTokeniWithParameterInPromptLine(t *testing.T) {
	prompt := PromptLine{"{time|yy/mm} {user}@{host}"}
	token := Token("time|yy/mm")

	prompt.Apply(token, "05/1992")

	assert.Equal(t, "05/1992 {user}@{host}", prompt.Text)
}
