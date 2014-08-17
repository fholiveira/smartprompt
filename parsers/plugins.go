package parsers

import (
	. "github.com/fholiveira/smartprompt/plugins"
	"github.com/fholiveira/smartprompt/plugins/location"
	"github.com/fholiveira/smartprompt/plugins/shell"
)

type PluginParser struct{}

func mapPlugins() map[string]Plugin {
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
	plugins := mapPlugins()

	errors := make([]error, 0)
	for _, token := range prompt.Tokens() {
		plugin, isPlugin := plugins[token.Name()]
		if !isPlugin {
			continue
		}

		pluginPrompt, err := plugin.Prompt(token.Parameter())
		if nil != err {
			errors = append(errors, err)
		}

		prompt.Apply(token, pluginPrompt)
	}

	if len(errors) > 0 {
		return PromptLine{prompt.Text}, errors
	}

	return PromptLine{prompt.Text}, nil
}
