package shell

type Release struct{}

func (release Release) Prompt() (string, error) {
	return "\\V", nil
}
