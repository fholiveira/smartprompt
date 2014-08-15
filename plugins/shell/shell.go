package shell

type Shell struct{}

func (shell Shell) Prompt() (string, error) {
	return "\\s", nil
}
