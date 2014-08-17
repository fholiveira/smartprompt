package parsers

import (
	. "github.com/fholiveira/smartprompt/plugins"
	"github.com/fholiveira/smartprompt/plugins/location"
	"github.com/fholiveira/smartprompt/plugins/shell"
)

type PluginParser struct{}

func (parser PluginParser) Plugins() map[string]Plugin {
	return map[string]Plugin{
		"user":              User{},
		"host":              Host{},
		"time":              DateTime{},
		"dir":               Directory{},
		"fqdn":              FullQualifiedDomainName{},
		"git":               Git{},
		"prompt:symbol":     PromptSymbol{},
		"location":          location.Default{},
		"shell":             shell.Shell{},
		"shell:version":     shell.Version{},
		"shell:release":     shell.Release{},
		"location:vimstyle": location.VimStyle{},
	}
}

func (parser PluginParser) Parse(prompt PromptLine) (PromptLine, []error) {
	plugins := parser.Plugins()

	errors := ErrorList{}.Init()
	for _, token := range prompt.Tokens() {
		plugin, isPlugin := plugins[token.Name()]
		if isPlugin {
			pluginPrompt, err := plugin.Prompt(token.Parameter())

			errors.Append(err)
			prompt.Apply(token, pluginPrompt)
		}
	}

	return PromptLine{prompt.Text}, errors.Items()
}
