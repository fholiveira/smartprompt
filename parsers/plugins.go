package parsers

import (
	. "github.com/fholiveira/smartprompt/plugins"
	"github.com/fholiveira/smartprompt/plugins/git"
	"github.com/fholiveira/smartprompt/plugins/location"
	"github.com/fholiveira/smartprompt/plugins/shell"
)

type PluginParser struct{}

var Plugins = func() map[string]Plugin {
	return map[string]Plugin{
		"user":          User{},
		"host":          Host{},
		"time":          DateTime{},
		"dir":           Directory{},
		"fqdn":          FullQualifiedDomainName{},
		"line:break":    LineBreak{},
		"prompt:symbol": PromptSymbol{},

		"virtualenv": Virtualenv{},

		"git":           git.GitStatus{},
		"sourcecontrol": SourceControl{},

		"shell":         shell.Shell{},
		"shell:version": shell.Version{},
		"shell:release": shell.Release{},

		"location":          location.Default{},
		"location:vimstyle": location.VimStyle{},
	}
}

func (parser PluginParser) Parse(prompt PromptLine) (PromptLine, []error) {
	plugins := Plugins()

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
