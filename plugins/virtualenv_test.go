package plugins

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenThereIsVirtualenv(t *testing.T) {
	getVirtualenv = func() string { return "/home/user/env" }
	venv, err := Virtualenv{}.Prompt(nil)

	if assert.NoError(t, err) {
		assert.Equal(t, "env", venv)
	}
}

func TestWhenThereIsNoVirtualenv(t *testing.T) {
	getVirtualenv = func() string { return "" }
	venv, err := Virtualenv{}.Prompt(nil)

	if assert.NoError(t, err) {
		assert.Equal(t, "", venv)
	}
}

func TestWhenVirtualenvIsAHiddenDirectory(t *testing.T) {
	getVirtualenv = func() string { return "/home/user/.env" }
	venv, err := Virtualenv{}.Prompt(nil)

	if assert.NoError(t, err) {
		assert.Equal(t, "env", venv)
	}
}

func TestWhenThereIsNoVirtualenvAndUserPassedParameters(t *testing.T) {
	getVirtualenv = func() string { return "" }
	venv, err := Virtualenv{}.Prompt([]string{"(", ")"})

	if assert.NoError(t, err) {
		assert.Equal(t, "", venv)
	}
}

func TestWhenThereIsNoVirtualenvAndUserPassedOneParameter(t *testing.T) {
	getVirtualenv = func() string { return "" }
	venv, err := Virtualenv{}.Prompt([]string{"->"})

	if assert.NoError(t, err) {
		assert.Equal(t, "", venv)
	}
}

func TestWhenUserPassParameters(t *testing.T) {
	getVirtualenv = func() string { return "/home/user/env" }
	venv, err := Virtualenv{}.Prompt([]string{"[", "]"})

	if assert.NoError(t, err) {
		assert.Equal(t, "[env]", venv)
	}
}

func TestWhenUserPassOnlyOneParameter(t *testing.T) {
	getVirtualenv = func() string { return "/home/user/env" }
	venv, err := Virtualenv{}.Prompt([]string{"-> "})

	if assert.NoError(t, err) {
		assert.Equal(t, "-> env", venv)
	}
}
