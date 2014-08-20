package parsers

import (
	"errors"
	"testing"

	. "github.com/fholiveira/smartprompt/plugins"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type PluginMock struct{ mock.Mock }

func (plugin PluginMock) Prompt(parameter string) (string, error) {
	args := plugin.Mock.Called(parameter)
	return args.String(0), args.Error(1)
}

type PluginParserTestSuite struct {
	suite.Suite
	parser PluginParser
}

func (suite *PluginParserTestSuite) SetupTest() {
	suite.parser = PluginParser{}

	common, parameter, error1, error2 := PluginMock{}, PluginMock{}, PluginMock{}, PluginMock{}

	common.On("Prompt", "").Return("common_plugin_value", nil)
	parameter.On("Prompt", "123").Return("p_value__123", nil)
	error1.On("Prompt", "").Return("", errors.New("Error 1"))
	error2.On("Prompt", "").Return("", errors.New("Error 2"))

	Plugins = func() map[string]Plugin {
		return map[string]Plugin{
			"common":    common,
			"parameter": parameter,
			"error1":    error1,
			"error2":    error2,
		}
	}
}

func (suite *PluginParserTestSuite) TestParsePlugin() {
	prompt := PromptLine{"{RED}{common}"}

	parsedPrompt, _ := suite.parser.Parse(prompt)
	assert.Equal(suite.T(), "{RED}common_plugin_value", parsedPrompt.Text)
}

func (suite *PluginParserTestSuite) TestShouldRaiseErrorOfPlugin() {
	prompt := PromptLine{"{YELLOW}{error1}"}

	_, errors := suite.parser.Parse(prompt)

	if assert.NotNil(suite.T(), errors) {
		assert.Equal(suite.T(), "Error 1", errors[0].Error())
	}
}

func (suite *PluginParserTestSuite) TestShouldRaiseErrorsOfMultiplePlugins() {
	prompt := PromptLine{"{CYAN}{error2}{BLACK}{error1}"}

	_, errors := suite.parser.Parse(prompt)

	if assert.NotNil(suite.T(), errors) {
		assert.Equal(suite.T(), "Error 2", errors[0].Error())
		assert.Equal(suite.T(), "Error 1", errors[1].Error())
	}
}

func (suite *PluginParserTestSuite) TestShouldRemoveTokenOfPluginWithError() {
	prompt := PromptLine{"{YELLOW}{error1}"}

	parsedPrompt, _ := suite.parser.Parse(prompt)
	assert.Equal(suite.T(), "{YELLOW}", parsedPrompt.Text)
}

func (suite *PluginParserTestSuite) TestShouldRemoveTokensOfMultiplePluginsWithError() {
	prompt := PromptLine{"{CYAN}{error1}{BLACK}{error2}"}

	parsedPrompt, _ := suite.parser.Parse(prompt)
	assert.Equal(suite.T(), "{CYAN}{BLACK}", parsedPrompt.Text)
}

func (suite *PluginParserTestSuite) TestParsePluginWithParameter() {
	prompt := PromptLine{"{BLUE}{parameter|123}"}

	parsedPrompt, _ := suite.parser.Parse(prompt)
	assert.Equal(suite.T(), "{BLUE}p_value__123", parsedPrompt.Text)
}

func TestPluginParserTestSuite(t *testing.T) {
	suite.Run(t, new(PluginParserTestSuite))
}
