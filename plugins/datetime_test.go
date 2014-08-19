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
	prompt, err := DateTime{}.Prompt("yy")
	assert := assert.New(suite.T())

	if assert.NoError(err) {
		assert.Equal("1992", prompt)
	}
}

func (suite *DateTimeTestSuite) TestLongMonthFormat() {
	prompt, err := DateTime{}.Prompt("mm")
	assert := assert.New(suite.T())

	if assert.NoError(err) {
		assert.Equal("05", prompt)
	}
}

func (suite *DateTimeTestSuite) TestLongDayFormat() {
	prompt, err := DateTime{}.Prompt("dd")
	assert := assert.New(suite.T())

	if assert.NoError(err) {
		assert.Equal("09", prompt)
	}
}

func (suite *DateTimeTestSuite) TestLongHourFormat() {
	prompt, err := DateTime{}.Prompt("hh")
	assert := assert.New(suite.T())

	if assert.NoError(err) {
		assert.Equal("16", prompt)
	}
}

func (suite *DateTimeTestSuite) TestLongMinuteFormat() {
	prompt, err := DateTime{}.Prompt("MM")
	assert := assert.New(suite.T())

	if assert.NoError(err) {
		assert.Equal("05", prompt)
	}
}

func (suite *DateTimeTestSuite) TestLongSecondFormat() {
	prompt, err := DateTime{}.Prompt("ss")
	assert := assert.New(suite.T())

	if assert.NoError(err) {
		assert.Equal("07", prompt)
	}
}

func (suite *DateTimeTestSuite) TestShortYearFormat() {
	prompt, err := DateTime{}.Prompt("y")
	assert := assert.New(suite.T())

	if assert.NoError(err) {
		assert.Equal("92", prompt)
	}
}

func (suite *DateTimeTestSuite) TestShortMonthFormat() {
	prompt, err := DateTime{}.Prompt("m")
	assert := assert.New(suite.T())

	if assert.NoError(err) {
		assert.Equal("5", prompt)
	}
}

func (suite *DateTimeTestSuite) TestShortDayFormat() {
	prompt, err := DateTime{}.Prompt("d")
	assert := assert.New(suite.T())

	if assert.NoError(err) {
		assert.Equal("9", prompt)
	}
}

func (suite *DateTimeTestSuite) TestShortHourFormat() {
	prompt, err := DateTime{}.Prompt("h")
	assert := assert.New(suite.T())

	if assert.NoError(err) {
		assert.Equal("04", prompt)
	}
}

func (suite *DateTimeTestSuite) TestShortMinuteFormat() {
	prompt, err := DateTime{}.Prompt("M")
	assert := assert.New(suite.T())

	if assert.NoError(err) {
		assert.Equal("5", prompt)
	}
}

func (suite *DateTimeTestSuite) TestShortSecondFormat() {
	prompt, err := DateTime{}.Prompt("s")
	assert := assert.New(suite.T())

	if assert.NoError(err) {
		assert.Equal("7", prompt)
	}
}

func TestDateTimeTestSuite(t *testing.T) {
	suite.Run(t, new(DateTimeTestSuite))
}
