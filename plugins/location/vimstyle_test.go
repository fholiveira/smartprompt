package location

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVimStyleLocationWhenCurrentDirIsRootDir(t *testing.T) {
	currentDirectory = func() (string, string, error) {
		return "/", "/", nil
	}

	location, err := VimStyle{}.Prompt(nil)
	if assert.NoError(t, err) {
		assert.Equal(t, "/", location)
	}
}

func TestVimStyleLocationWithSymlinks(t *testing.T) {
	currentDirectory = func() (string, string, error) {
		return "~/Documents/2014/papers", "~/docs/papers", nil
	}

	location, err := VimStyle{}.Prompt(nil)

	if assert.NoError(t, err) {
		assert.Equal(t, "~/d/papers", location)
	}
}

func TestVimStyleAbsoluteLocation(t *testing.T) {
	currentDirectory = func() (string, string, error) {
		return "~/Documents/2014/papers", "~/papers", nil
	}

	location, err := VimStyle{}.Prompt([]string{"absolute"})

	if assert.NoError(t, err) {
		assert.Equal(t, "~/D/2/papers", location)
	}
}

func TestVimStyleLocationWhenCurrentDirIsHomeDir(t *testing.T) {
	currentDirectory = func() (string, string, error) {
		return "~", "~", nil
	}

	location, err := VimStyle{}.Prompt(nil)
	if assert.NoError(t, err) {
		assert.Equal(t, "~", location)
	}
}

func TestVimStyleLocationWhenCurrentDirIsiChildOfRootDir(t *testing.T) {
	currentDirectory = func() (string, string, error) {
		return "/home/Documents/20140322/test", "/test", nil
	}

	location, err := VimStyle{}.Prompt([]string{"absolute"})
	if assert.NoError(t, err) {
		assert.Equal(t, "/h/D/2/test", location)
	}
}

func TestVimStyleLocationWhenCurrentDirIsChildOfHomeDir(t *testing.T) {
	currentDirectory = func() (string, string, error) {
		return "~/Projects/golang/src/smartprompt", "~/smartprompt", nil
	}

	location, err := VimStyle{}.Prompt([]string{"absolute"})
	if assert.NoError(t, err) {
		assert.Equal(t, "~/P/g/s/smartprompt", location)
	}
}
