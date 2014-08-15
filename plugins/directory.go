package plugins

type Directory struct{}

func (dir Directory) Prompt(parameter string) (string, error) {
	return "\\W", nil
}
