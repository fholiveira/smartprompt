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

func TestVimStyleAbsoluteLocationWithHiddenDir(t *testing.T) {
	currentDirectory = func() (string, string, error) {
		return "/home/username/docs/.math/papers", "~/docs/.math/papers", nil
	}

	location, err := VimStyle{}.Prompt([]string{"absolute"})

	if assert.NoError(t, err) {
		assert.Equal(t, "/h/u/d/.m/papers", location)
	}
}

func TestVimStyleLocationWithHiddenDir(t *testing.T) {
	currentDirectory = func() (string, string, error) {
		return "/home/username/docs/.math/papers", "~/docs/.math/papers", nil
	}

	location, err := VimStyle{}.Prompt(nil)

	if assert.NoError(t, err) {
		assert.Equal(t, "~/d/.m/papers", location)
	}
}

func TestVimStyleLocationWithHiddenSymbolicDir(t *testing.T) {
	currentDirectory = func() (string, string, error) {
		return "/home/username/docs/.math/papers", "~/docs/.math/papers", nil
	}

	isSymlink = func(path string) bool {
		return path == "~/docs/.math"
	}

	location, err := VimStyle{}.Prompt(nil)

	if assert.NoError(t, err) {
		assert.Equal(t, "~/d/.m@/papers", location)
	}
}

func TestVimStyleLocationWithSymlinks(t *testing.T) {
	currentDirectory = func() (string, string, error) {
		return "/home/username/Documents/2014/papers", "~/docs/papers", nil
	}

	location, err := VimStyle{}.Prompt(nil)

	if assert.NoError(t, err) {
		assert.Equal(t, "~/d/papers", location)
	}
}

func TestVimStyleAbsoluteLocation(t *testing.T) {
	currentDirectory = func() (string, string, error) {
		return "/home/username/Documents/2014/papers",
			"~/Documents/2014/papers",
			nil
	}

	location, err := VimStyle{}.Prompt([]string{"absolute"})

	if assert.NoError(t, err) {
		assert.Equal(t, "/h/u/D/2/papers", location)
	}
}

func TestVimStyleLocationWhenCurrentDirIsHomeDir(t *testing.T) {
	currentDirectory = func() (string, string, error) {
		return "/home/username", "~", nil
	}

	location, err := VimStyle{}.Prompt(nil)
	if assert.NoError(t, err) {
		assert.Equal(t, "~", location)
	}
}

func TestVimStyleLocationWhenCurrentDirIsiChildOfRootDir(t *testing.T) {
	currentDirectory = func() (string, string, error) {
		return "/Documents/20140322/test", "/Documents/20140322/test", nil
	}

	location, err := VimStyle{}.Prompt([]string{"absolute"})
	if assert.NoError(t, err) {
		assert.Equal(t, "/D/2/test", location)
	}
}

func TestVimStyleLocationWhenCurrentDirIsChildOfHomeDir(t *testing.T) {
	currentDirectory = func() (string, string, error) {
		return "/home/username/Projects/golang/src/smartprompt",
			"~/Projects/golang/src/smartprompt",
			nil
	}

	location, err := VimStyle{}.Prompt([]string{"absolute"})
	if assert.NoError(t, err) {
		assert.Equal(t, "/h/u/P/g/s/smartprompt", location)
	}
}
