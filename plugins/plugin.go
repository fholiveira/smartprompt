package plugins

type Plugin interface {
	Prompt(parameters []string) (string, error)
}
