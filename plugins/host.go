package plugins

type Host struct{}

func (host Host) Prompt(parameter string) (string, error) {
	return "\\h", nil
}
