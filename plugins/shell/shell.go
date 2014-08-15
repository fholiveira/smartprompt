package shell

type Shell struct{}

func (shell Shell) Prompt(parameter string) (string, error) {
	return "\\s", nil
}
