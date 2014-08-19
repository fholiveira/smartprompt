package parsers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTokenNameWhenTokenHasParameter(t *testing.T) {
	name := Token("name123|parameter987").Name()
	assert.Equal(t, "name123", name)
}

func TestGetTokenNameWhenTokenDoesNotHaveParameter(t *testing.T) {
	name := Token("name123").Name()
	assert.Equal(t, "name123", name)
}

func TestGetTokenParameterWhenTokenHasParameter(t *testing.T) {
	parameter := Token("name123|parameter987").Parameter()
	assert.Equal(t, "parameter987", parameter)
}

func TestGetTokenParameterWhenTokenDoesNotHaveParameter(t *testing.T) {
	parameter := Token("name123").Parameter()
	assert.Empty(t, parameter)
}
