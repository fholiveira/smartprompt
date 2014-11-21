package plugins

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectory(t *testing.T) {
	directory, err := Directory{}.Prompt(nil)

	if assert.NoError(t, err) {
		assert.Equal(t, "\\W", directory)
	}
}
