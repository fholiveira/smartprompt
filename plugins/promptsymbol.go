package plugins

import "syscall"

type PromptSymbol struct{}

func (promptSymbol PromptSymbol) Prompt(parameter string) (string, error) {
	symbol := "$"

	if syscall.Geteuid() == 0 {
		symbol = "#"
	}

	return symbol, nil
}
