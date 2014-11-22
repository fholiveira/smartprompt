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

func TestCustomPromptSymbolWhenUserIsRoot(t *testing.T) {
	userIsRoot = func() bool { return true }
	symbol, err := PromptSymbol{}.Prompt([]string{"<>", "%%"})

	if assert.NoError(t, err) {
		assert.Equal(t, "<>", symbol)
	}
}

func TestCustomPromptSymbolWhenUserIsNotRoot(t *testing.T) {
	userIsRoot = func() bool { return false }
	symbol, err := PromptSymbol{}.Prompt([]string{">>", "->"})

	if assert.NoError(t, err) {
		assert.Equal(t, "->", symbol)
	}
}

func TestCustomPromptSymbolWhenUserIsRootAndUserPassOnlyOneParameter(t *testing.T) {
	userIsRoot = func() bool { return true }
	symbol, err := PromptSymbol{}.Prompt([]string{"<>"})

	if assert.NoError(t, err) {
		assert.Equal(t, "<>", symbol)
	}
}

func TestCustomPromptSymbolWhenUserIsNotRootAndUserPassOnlyOneParameter(t *testing.T) {
	userIsRoot = func() bool { return false }
	symbol, err := PromptSymbol{}.Prompt([]string{">>"})

	if assert.NoError(t, err) {
		assert.Equal(t, "$", symbol)
	}
}
