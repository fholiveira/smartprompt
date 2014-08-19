package parsers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColorParserShouldParseColor(t *testing.T) {
	prompt := PromptLine{"{user}@{RED}{host}"}

	parsedPrompt, _ := ColorParser{}.Parse(prompt)
	assert.Equal(t, "{user}@\\e[0;31m{host}", parsedPrompt.Text)
}

func TestColorParserShouldParseColorUnderline(t *testing.T) {
	prompt := PromptLine{"{user}@{BLUE:underline}{host}"}

	parsedPrompt, _ := ColorParser{}.Parse(prompt)
	assert.Equal(t, "{user}@\\e[4;34m{host}", parsedPrompt.Text)
}

func TestColorParserShouldParseColorBold(t *testing.T) {
	prompt := PromptLine{"{user}@{CYAN:bold}{host}"}

	parsedPrompt, _ := ColorParser{}.Parse(prompt)
	assert.Equal(t, "{user}@\\e[1;36m{host}", parsedPrompt.Text)
}

func TestColorParserShouldParseColorBackground(t *testing.T) {
	prompt := PromptLine{"{user}@{YELLOW:background}{host}"}

	parsedPrompt, _ := ColorParser{}.Parse(prompt)
	assert.Equal(t, "{user}@\\e[43m{host}", parsedPrompt.Text)
}

func TestColorParserShouldParseTextReset(t *testing.T) {
	prompt := PromptLine{"{user}@{TEXT:reset}{host}"}

	parsedPrompt, _ := ColorParser{}.Parse(prompt)
	assert.Equal(t, "{user}@\\e[0m{host}", parsedPrompt.Text)
}
