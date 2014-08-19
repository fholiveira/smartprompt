package plugins

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHost(t *testing.T) {
	host, err := Host{}.Prompt("")

	if assert.NoError(t, err) {
		assert.Equal(t, "\\h", host)
	}
}
