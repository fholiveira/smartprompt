package plugins

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromptSymbolWhenUserIsRoot(t *testing.T) {
	userIsRoot = func() bool { return true }
	symbol, err := PromptSymbol{}.Prompt(nil)

	if assert.NoError(t, err) {
		assert.Equal(t, "#", symbol)
	}
}

func TestPromptSymbolWhenUserIsNotRoot(t *testing.T) {
	userIsRoot = func() bool { return false }
	symbol, err := PromptSymbol{}.Prompt(nil)

	if assert.NoError(t, err) {
		assert.Equal(t, "$", symbol)
	}
}
