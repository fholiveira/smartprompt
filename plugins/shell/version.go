package shell

type Version struct{}

func (version Version) Prompt(parameter string) (string, error) {
	return "\\v", nil
}
