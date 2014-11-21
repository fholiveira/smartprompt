package shell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShellRelease(t *testing.T) {
	release, err := Release{}.Prompt(nil)

	if assert.NoError(t, err) {
		assert.Equal(t, "\\V", release)
	}
}
