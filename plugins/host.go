package plugins

type Host struct{}

func (host Host) Prompt() (string, error) {
	return "\\h", nil
}
