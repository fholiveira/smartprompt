package parsers

import (
	"strings"

	. "github.com/fholiveira/smartprompt/plugins"
	"github.com/fholiveira/smartprompt/plugins/location"
	"github.com/fholiveira/smartprompt/plugins/shell"
)

type PluginParser struct{}

func mapPlugins() map[string]Plugin {
	return map[string]Plugin{
		"{user}":              User{},
		"{host}":              Host{},
		"{dir}":               Directory{},
		"{fqdn}":              FullQualifiedDomainName{},
		"{git}":               Git{},
		"{prompt:symbol}":     PromptSymbol{},
		"{location}":          location.Default{},
		"{shell}":             shell.Shell{},
		"{shell:version}":     shell.Version{},
		"{shell:release}":     shell.Release{},
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

		pluginPrompt, err := plugin.Prompt()
		if nil != err {
			return "", err
		}

		prompt = strings.Replace(prompt, token, pluginPrompt, -1)
	}

	return prompt, nil
}
