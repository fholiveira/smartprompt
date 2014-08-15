package main

import (
	"fmt"

	"github.com/docopt/docopt-go"
	. "github.com/fholiveira/smartprompt/parsers"
)

func parsePrompt(pattern string, debug bool) string {
	parsers := []Parser{
		PluginParser{},
		ColorParser{},
		WhiteSpacesParser{},
	}

	var errors []error
	prompt := PromptLine{pattern}

	for _, parser := range parsers {
		prompt, errors = parser.Parse(prompt)
		if debug && nil != errors {
			for _, err := range errors {
				fmt.Println(err)
			}
		}
	}

	return prompt.Text
}

func main() {
	arguments, _ := docopt.Parse(usage(), nil, true, "0.1", false)

	debug, _ := arguments["--debug"].(bool)
	pattern, _ := arguments["--pattern"].(string)

	fmt.Println(parsePrompt(pattern, debug))
}

func usage() string {
	return `Smart Prompt
A customizable prompt generator for your terminal emulator.

Usage:
  smartprompt [--pattern=<pattern>] [-d | --debug]
  smartprompt -h | --help
  smartprompt --version

Options:
  --pattern=<pattern>    Prompt pattern [default: {GREEN:bold}{user}@{host} {BLUE:bold}{location:vimstyle} {git} {PURPLE:bold}{prompt:symbol} {TEXT:reset}].
  -d, --debug            Debug mode (print all errors in stdout).
  -h, --help             Show this screen.
  --version              Show version.`
}
