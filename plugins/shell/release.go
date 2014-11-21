package shell

type Release struct{}

func (release Release) Prompt(parameters []string) (string, error) {
	return "\\V", nil
}
