package shell

type Version struct{}

func (version Version) Prompt() (string, error) {
	return "\\v", nil
}
