package plugins

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLineBreak(t *testing.T) {
	host, err := LineBreak{}.Prompt(nil)

	if assert.NoError(t, err) {
		assert.Equal(t, "\\n", host)
	}
}
