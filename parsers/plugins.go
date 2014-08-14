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
		"{location}":          location.Default{},
		"{location:vimstyle}": location.VimStyle{},
	}
}

func (parser PluginParser) Parse(prompt string) (string, error) {
	for key, plugin := range mapPlugins() {
		if !strings.Contains(prompt, key) {
			continue
		}

		pluginPrompt, err := plugin.Prompt()
		if nil == err {
			prompt = strings.Replace(prompt, key, pluginPrompt, -1)
		}

	}

	return prompt, nil
}
