package plugins

type Host struct{}

func (host Host) Prompt(parameters []string) (string, error) {
	return "\\h", nil
}
