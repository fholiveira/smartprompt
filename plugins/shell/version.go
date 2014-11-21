package shell

type Version struct{}

func (version Version) Prompt(parameters []string) (string, error) {
	return "\\v", nil
}
