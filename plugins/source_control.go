package plugins

import "github.com/fholiveira/smartprompt/plugins/git"

type SourceControl struct{}

type SourceControlPlugin interface {
	IsApplicable() bool
}

func (sourceControl SourceControl) Prompt(parameter string) (string, error) {
	plugins := [...]SourceControlPlugin{
		git.GitStatus{},
	}

	for _, plugin := range plugins {
		if plugin.IsApplicable() {
			return (plugin.(Plugin)).Prompt(parameter)
		}
	}

	return "", nil
}