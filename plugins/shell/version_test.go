package shell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShellVersion(t *testing.T) {
	version, err := Version{}.Prompt(nil)

	if assert.NoError(t, err) {
		assert.Equal(t, "\\v", version)
	}
}
