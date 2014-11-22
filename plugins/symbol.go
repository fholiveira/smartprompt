package plugins

import "syscall"

type PromptSymbol struct{}

var userIsRoot = func() bool { return syscall.Geteuid() == 0 }

func (plugin PromptSymbol) Prompt(parameters []string) (string, error) {
	root, common := plugin.load_symbols(parameters)

	if userIsRoot() {
		return root, nil
	}

	return common, nil
}

func (plugin PromptSymbol) load_symbols(parameters []string) (string, string) {
	root, common := "#", "$"
	if len(parameters) == 0 {
		return root, common
	}

	if len(parameters) == 1 {
		return parameters[0], common
	}

	return parameters[0], parameters[1]
}
