package parsers

import (
	"strings"

	. "github.com/fholiveira/smartprompt/plugins"
	"github.com/fholiveira/smartprompt/plugins/location"
)

type PluginParser struct{}

func mapPlugins() map[string]Plugin {
	return map[string]Plugin{
		"{user}":              User{},
		"{host}":              Host{},
		"{git}":               Git{},
		"{prompt:symbol}":     PromptSymbol{},
		"{location}":          location.Default{},
		"{location:vimstyle}": location.VimStyle{},
	}
}

func (parser PluginParser) Parse(prompt string) (string, error) {
	plugins := mapPlugins()

	for _, token := range getTokens(prompt) {
		plugin, isPlugin := plugins[token]
		if !isPlugin {
			continue
		}

		pluginPrompt, _ := plugin.Prompt()
		prompt = strings.Replace(prompt, token, pluginPrompt, -1)
	}

	return prompt, nil
}
