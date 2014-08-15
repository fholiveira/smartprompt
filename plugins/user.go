package plugins

type User struct{}

func (user User) Prompt(parameter string) (string, error) {
	return "\\u", nil
}
