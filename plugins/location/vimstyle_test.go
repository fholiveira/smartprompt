package location

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVimStyleLocationWhenCurrentDirIsRootDir(t *testing.T) {
	getWorkingDir = func() (string, error) { return "/", nil }
	location, err := VimStyle{}.Prompt("")

	if assert.NoError(t, err) {
		assert.Equal(t, "/", location)
	}
}

func TestVimStyleLocationWhenCurrentDirIsHomeDir(t *testing.T) {
	getWorkingDir = func() (string, error) { return "~", nil }
	location, err := VimStyle{}.Prompt("")

	if assert.NoError(t, err) {
		assert.Equal(t, "~", location)
	}
}

func TestVimStyleLocationWhenCurrentDirIsiChildOfRootDir(t *testing.T) {
	getWorkingDir = func() (string, error) { return "/home/Documents/20140322/test", nil }
	location, err := VimStyle{}.Prompt("")

	if assert.NoError(t, err) {
		assert.Equal(t, "/h/D/2/test", location)
	}
}

func TestVimStyleLocationWhenCurrentDirIsChildOfHomeDir(t *testing.T) {
	getWorkingDir = func() (string, error) { return "~/Projects/golang/src/smartprompt", nil }
	location, err := VimStyle{}.Prompt("")

	if assert.NoError(t, err) {
		assert.Equal(t, "~/P/g/s/smartprompt", location)
	}
}
