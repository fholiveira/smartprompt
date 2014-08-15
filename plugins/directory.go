package plugins

type Directory struct{}

func (dir Directory) Prompt() (string, error) {
	return "\\W", nil
}
