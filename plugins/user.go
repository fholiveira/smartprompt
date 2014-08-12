package plugins

type User struct{}

func (user User) Prompt() (string, error) {
	return "\\u", nil
}
