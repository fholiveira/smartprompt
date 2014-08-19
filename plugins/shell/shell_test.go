package shell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShell(t *testing.T) {
	shell, err := Shell{}.Prompt("")

	if assert.NoError(t, err) {
		assert.Equal(t, "\\s", shell)
	}
}
