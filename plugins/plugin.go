package plugins

type Plugin interface {
	Prompt() (string, error)
}
