package shell

type Shell struct{}

func (shell Shell) Prompt(parameters []string) (string, error) {
	return "\\s", nil
}
