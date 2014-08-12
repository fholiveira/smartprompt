package location

type Default struct{}

func (location Default) Prompt() (string, error) {
	return "\\h", nil
}
