package parsers

import (
	"errors"
	"testing"

	. "github.com/fholiveira/smartprompt/plugins"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CommonPlugin struct{}

func (plugin CommonPlugin) Prompt(parameter string) (string, error) {
	return "common_plugin_value", nil
}

type PluginWithParameter struct{}

func (plugin PluginWithParameter) Prompt(parameter string) (string, error) {
	return "p_value__" + parameter, nil
}

type PluginWithError1 struct{}

func (plugin PluginWithError1) Prompt(parameter string) (string, error) {
	return "", errors.New("Error 1")
}

type PluginWithError2 struct{}

func (plugin PluginWithError2) Prompt(parameter string) (string, error) {
	return "", errors.New("Error 2")
}

type PluginParserTestSuite struct {
	suite.Suite
	parser PluginParser
}

func (suite *PluginParserTestSuite) SetupTest() {
	suite.parser = PluginParser{}

	Plugins = func() map[string]Plugin {
		return map[string]Plugin{
			"common":    CommonPlugin{},
			"parameter": PluginWithParameter{},
			"error1":    PluginWithError1{},
			"error2":    PluginWithError2{},
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
