package plugins

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullQualifiedDomainName(t *testing.T) {
	fqdn, err := FullQualifiedDomainName{}.Prompt("")

	if assert.NoError(t, err) {
		assert.Equal(t, "\\H", fqdn)
	}
}
