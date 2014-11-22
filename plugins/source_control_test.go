package plugins

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type SourceControlPluginMock struct{ mock.Mock }

func (plugin SourceControlPluginMock) Prompt(parameters []string) (string, error) {
	args := plugin.Mock.Called(parameters)
	return args.String(0), args.Error(1)
}

func (plugin SourceControlPluginMock) IsApplicable() bool {
	args := plugin.Mock.Called()
	return args.Bool(0)
}

type SourceControlTestSuite struct {
	suite.Suite
}

func (suite *SourceControlTestSuite) TestShoulNotApplyIfNotApplicable() {
	git := SourceControlPluginMock{}
	git.On("IsApplicable").Return(false)

	plugins = func() []SourceControlPlugin {
		return []SourceControlPlugin{git}
	}

	prompt, _ := SourceControl{}.Prompt([]string{})

	assert.Equal(suite.T(), "", prompt)
}

func (suite *SourceControlTestSuite) TestShoulApplyOnlyOnePLugin() {
	git := SourceControlPluginMock{}
	git.On("IsApplicable").Return(true)
	git.On("Prompt", []string{}).Return("GIT", nil)

	svn := SourceControlPluginMock{}
	svn.On("IsApplicable").Return(true)
	svn.On("Prompt", []string{}).Return("SVN", nil)

	plugins = func() []SourceControlPlugin {
		return []SourceControlPlugin{git, svn}
	}

	prompt, _ := SourceControl{}.Prompt([]string{})

	assert.Equal(suite.T(), "GIT", prompt)
}

func TestSourceControlTestSuite(t *testing.T) {
	suite.Run(t, new(SourceControlTestSuite))
}
