package plugins

type LineBreak struct{}

func (line LineBreak) Prompt(parameters []string) (string, error) {
	return "\\n", nil
}
