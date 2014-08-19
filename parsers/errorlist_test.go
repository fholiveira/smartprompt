package parsers

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorListShouldInitAnEmptyErrosArray(t *testing.T) {
	list := ErrorList{}.Init()

	assert.Len(t, list.errors, 0)
}

func TestErrorListItemsShouldReturnNilIfItDoesNotHaveErrors(t *testing.T) {
	list := ErrorList{}.Init()

	assert.Nil(t, list.Items())
}

func TestErrorListShouldReturnAnArrayOfErrors(t *testing.T) {
	errs := []error{errors.New("Error 1"), errors.New("Error 2")}
	list := ErrorList{}.Init()

	list.Append(errs[0])
	list.Append(errs[1])

	assert.Equal(t, errs, list.errors)
}

func TestErrorListShouldAddError(t *testing.T) {
	err := errors.New("Error 1")
	list := ErrorList{}.Init()

	list.Append(err)

	assert.Equal(t, err, list.errors[0])
}

func TestErrorListShouldNotAddNilError(t *testing.T) {
	list := ErrorList{}.Init()

	list.Append(nil)

	assert.Nil(t, list.Items())
}
