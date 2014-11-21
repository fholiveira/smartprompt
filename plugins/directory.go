package plugins

type Directory struct{}

func (dir Directory) Prompt(parameters []string) (string, error) {
	return "\\W", nil
}
