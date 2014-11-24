package location

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSymlinkLocation(t *testing.T) {
	currentDirectory = func() (string, string, error) {
		return "~/Documents/2014/papers", "~/papers", nil
	}

	location, err := Default{}.Prompt(nil)

	if assert.NoError(t, err) {
		assert.Equal(t, "~/papers", location)
	}
}

func TestAbsoluteLocation(t *testing.T) {
	currentDirectory = func() (string, string, error) {
		return "~/Documents/2014/papers", "~/papers", nil
	}

	location, err := Default{}.Prompt([]string{"absolute"})

	if assert.NoError(t, err) {
		assert.Equal(t, "~/Documents/2014/papers", location)
	}
}
