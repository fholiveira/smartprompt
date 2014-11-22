package plugins

import "github.com/fholiveira/smartprompt/plugins/git"

type SourceControlPlugin interface {
	IsApplicable() bool
}

type SourceControl struct{}

var plugins = func() []SourceControlPlugin {
	return []SourceControlPlugin{
		git.GitStatus{},
	}
}

func (sourceControl SourceControl) Prompt(parameters []string) (string, error) {
	for _, plugin := range plugins() {
		if plugin.IsApplicable() {
			return (plugin.(Plugin)).Prompt(parameters)
		}
	}

	return "", nil
}
