package plugins

type LineBreak struct{}

func (line LineBreak) Prompt(parameter string) (string, error) {
	return "\\n", nil
}
