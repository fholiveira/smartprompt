package parsers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhiteSpacesParser(t *testing.T) {
	promptLine := PromptLine{"{RED}  {BLUE}   {CYAN}    {GREEN}"}

	parsedPrompt, err := WhiteSpacesParser{}.Parse(promptLine)

	if assert.Nil(t, err) {
		assert.Equal(t, parsedPrompt.Text, "{RED} {BLUE} {CYAN} {GREEN}")
	}
}
