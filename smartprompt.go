package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)
import (
	. "github.com/fholiveira/smartprompt/plugins"
	"github.com/fholiveira/smartprompt/plugins/location"
)

func mapPlugins() map[string]Plugin {
	return map[string]Plugin{
		"{user}":              User{},
		"{host}":              Host{},
		"{git}":               Git{},
		"{location}":          location.Default{},
		"{location:vimstyle}": location.VimStyle{},
	}
}

func loadPromptPattern() (string, error) {
	args := os.Args[1:]

	if len(args) != 1 {
		return "", errors.New("Invalid arguments")
	}

	return args[0], nil
}

func createPrompt(pattern string, plugins map[string]Plugin) string {
	for key, plugin := range plugins {
		if !strings.Contains(pattern, key) {
			continue
		}

		prompt, err := plugin.Prompt()
		if err == nil {
			pattern = strings.Replace(pattern, key, prompt, -1)
		}

	}

	return pattern
}

func main() {
	promptPattern, err := loadPromptPattern()
	if nil != err {
		return
	}

	prompt := createPrompt(promptPattern, mapPlugins())
	fmt.Println(prompt)
}
