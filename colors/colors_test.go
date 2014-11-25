package colors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoldColor(t *testing.T) {
	var color Color = 15
	assert.Equal(t, "\\[\\e[1;15m\\]", color.Bold())
}

func TestNormalColor(t *testing.T) {
	var color Color = 15
	assert.Equal(t, "\\[\\e[0;15m\\]", color.Normal())
}

func TestUnderlineColor(t *testing.T) {
	var color Color = 15
	assert.Equal(t, "\\[\\e[4;15m\\]", color.Underline())
}

func TestBackgroundColor(t *testing.T) {
	var color Color = 15
	assert.Equal(t, "\\[\\e[15m\\]", color.Background())
}
