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
		"user":       User{},
		"host":       Host{},
		"time":       DateTime{},
		"dir":        Directory{},
		"fqdn":       FullQualifiedDomainName{},
		"line:break": LineBreak{},
		"symbol":     PromptSymbol{},

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

func validPlugins(tokens []Token) []InputMessage {
	plugins, messages := Plugins(), make([]InputMessage, 0)

	for _, token := range tokens {
		plugin, isPlugin := plugins[token.Name()]
		if isPlugin {
			messages = append(messages, InputMessage{plugin, token})
		}
	}

	return messages
}

func (parser PluginParser) Parse(prompt PromptLine) (PromptLine, []error) {
	errors, plugins := ErrorList{}.Init(), validPlugins(prompt.Tokens("|"))

	results := ParallelMap(plugins, func(message InputMessage) OutputMessage {
		result, err := message.plugin.Prompt(message.token.Parameters())
		return OutputMessage{message.token, result, err}
	})

	for _, message := range results {
		errors.Append(message.err)
		prompt.Apply(message.token, message.result)
	}

	return PromptLine{prompt.Text}, errors.Items()
}
