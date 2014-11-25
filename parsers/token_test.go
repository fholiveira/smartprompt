package parsers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTokenNameWhenTokenHasParameter(t *testing.T) {
	name := NewToken("name123|parameter987", "|").Name()
	assert.Equal(t, "name123", name)
}

func TestGetTokenNameWhenTokenDoesNotHaveParameter(t *testing.T) {
	name := NewToken("name123", "|").Name()
	assert.Equal(t, "name123", name)
}

func TestGetTokenParameterWhenTokenHasParameter(t *testing.T) {
	parameters := NewToken("name123|parameter987", "|").Parameters()
	assert.Equal(t, "parameter987", parameters[0])
}

func TestGetTokenParameterWhenTokenDoesNotHaveParameter(t *testing.T) {
	parameters := NewToken("name123", "|").Parameters()
	assert.Nil(t, parameters)
}
