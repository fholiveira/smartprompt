package plugins

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	user, err := User{}.Prompt(nil)

	if assert.NoError(t, err) {
		assert.Equal(t, "\\u", user)
	}
}
