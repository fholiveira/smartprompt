package plugins

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DateTimeTestSuite struct {
	suite.Suite
}

func (suite *DateTimeTestSuite) SetupTest() {
	now = func() time.Time {
		value, _ := time.Parse("02/01/2006 15:04:05", "09/05/1992 16:05:07")
		return value
	}
}

func (suite *DateTimeTestSuite) TestLongYearFormat() {
	prompt, err := DateTime{}.Prompt([]string{"yy"})
	assert := assert.New(suite.T())

	if assert.NoError(err) {
		assert.Equal("1992", prompt)
	}
}

func (suite *DateTimeTestSuite) TestLongMonthFormat() {
	prompt, err := DateTime{}.Prompt([]string{"mm"})
	assert := assert.New(suite.T())

	if assert.NoError(err) {
		assert.Equal("05", prompt)
	}
}

func (suite *DateTimeTestSuite) TestLongDayFormat() {
	prompt, err := DateTime{}.Prompt([]string{"dd"})
	assert := assert.New(suite.T())

	if assert.NoError(err) {
		assert.Equal("09", prompt)
	}
}

func (suite *DateTimeTestSuite) TestLongHourFormat() {
	prompt, err := DateTime{}.Prompt([]string{"hh"})
	assert := assert.New(suite.T())

	if assert.NoError(err) {
		assert.Equal("16", prompt)
	}
}

func (suite *DateTimeTestSuite) TestLongMinuteFormat() {
	prompt, err := DateTime{}.Prompt([]string{"MM"})
	assert := assert.New(suite.T())

	if assert.NoError(err) {
		assert.Equal("05", prompt)
	}
}

func (suite *DateTimeTestSuite) TestLongSecondFormat() {
	prompt, err := DateTime{}.Prompt([]string{"ss"})
	assert := assert.New(suite.T())

	if assert.NoError(err) {
		assert.Equal("07", prompt)
	}
}

func (suite *DateTimeTestSuite) TestShortYearFormat() {
	prompt, err := DateTime{}.Prompt([]string{"y"})
	assert := assert.New(suite.T())

	if assert.NoError(err) {
		assert.Equal("92", prompt)
	}
}

func (suite *DateTimeTestSuite) TestShortMonthFormat() {
	prompt, err := DateTime{}.Prompt([]string{"m"})
	assert := assert.New(suite.T())

	if assert.NoError(err) {
		assert.Equal("5", prompt)
	}
}

func (suite *DateTimeTestSuite) TestShortDayFormat() {
	prompt, err := DateTime{}.Prompt([]string{"d"})
	assert := assert.New(suite.T())

	if assert.NoError(err) {
		assert.Equal("9", prompt)
	}
}

func (suite *DateTimeTestSuite) TestShortHourFormat() {
	prompt, err := DateTime{}.Prompt([]string{"h"})
	assert := assert.New(suite.T())

	if assert.NoError(err) {
		assert.Equal("04", prompt)
	}
}

func (suite *DateTimeTestSuite) TestShortMinuteFormat() {
	prompt, err := DateTime{}.Prompt([]string{"M"})
	assert := assert.New(suite.T())

	if assert.NoError(err) {
		assert.Equal("5", prompt)
	}
}

func (suite *DateTimeTestSuite) TestShortSecondFormat() {
	prompt, err := DateTime{}.Prompt([]string{"s"})
	assert := assert.New(suite.T())

	if assert.NoError(err) {
		assert.Equal("7", prompt)
	}
}

func TestDateTimeTestSuite(t *testing.T) {
	suite.Run(t, new(DateTimeTestSuite))
}
