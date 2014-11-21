package plugins

type User struct{}

func (user User) Prompt(parameters []string) (string, error) {
	return "\\u", nil
}
