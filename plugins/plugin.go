package plugins

type Plugin interface {
	Prompt(parameter string) (string, error)
}
