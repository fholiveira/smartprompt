package main

import (
	"errors"
	"fmt"
	"os"

	. "github.com/fholiveira/smartprompt/parsers"
)

func loadPromptPattern() (string, error) {
	args := os.Args[1:]

	if len(args) != 1 {
		return "", errors.New("Invalid arguments")
	}

	return args[0], nil
}

func main() {
	promptPattern, err := loadPromptPattern()
	if nil != err {
		return
	}

	prompt, _ := ColorParser{}.Parse(promptPattern)
	prompt, _ = PluginParser{}.Parse(prompt)

	fmt.Println(prompt)
}
