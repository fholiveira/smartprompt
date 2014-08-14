package main

import (
	"fmt"
	"os"

	. "github.com/fholiveira/smartprompt/parsers"
)

func loadPromptPattern() string {
	args := os.Args[1:]

	if len(args) != 1 {
		return "{GREEN:bold}{user}@{host} {BLUE:bold}{location:vimstyle} {git} {CYAN:bold}{prompt:symbol} {TEXT:reset}"
	}

	return args[0]
}

func main() {
	prompt := loadPromptPattern()

	prompt, _ = PluginParser{}.Parse(prompt)
	prompt, _ = ColorParser{}.Parse(prompt)

	fmt.Println(prompt)
}
