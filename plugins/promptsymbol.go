package plugins

import "syscall"

type PromptSymbol struct{}

var userIsRoot = func() bool { return syscall.Geteuid() == 0 }

func (promptSymbol PromptSymbol) Prompt(parameters []string) (string, error) {
	symbol := "$"

	if userIsRoot() {
		symbol = "#"
	}

	return symbol, nil
}
