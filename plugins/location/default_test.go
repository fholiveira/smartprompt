package location

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultLocation(t *testing.T) {
	getWorkingDir = func() (string, error) { return "~/Documents", nil }
	location, err := Default{}.Prompt(nil)

	if assert.NoError(t, err) {
		assert.Equal(t, "~/Documents", location)
	}
}
