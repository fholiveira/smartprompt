package shell

type Release struct{}

func (release Release) Prompt(parameter string) (string, error) {
	return "\\V", nil
}
