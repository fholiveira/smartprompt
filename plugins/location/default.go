package location

import (
	"os"
	"os/user"
	"strings"
)

type Default struct{}

var currentDirectory = func() (string, string, error) {
	path, err := os.Getwd()
	if nil != err {
		return "", "", err
	}

	user, err := user.Current()
	if nil != err {
		return "", "", err
	}

	absolutePath := strings.Replace(path, user.HomeDir, "~", 1)
	symlinkPath := strings.Replace(os.Getenv("PWD"), user.HomeDir, "~", 1)

	return absolutePath, symlinkPath, nil
}

func (location Default) Prompt(parameters []string) (string, error) {
	absolutePath, symlinkPath, err := currentDirectory()
	path := symlinkPath

	if len(parameters) > 0 && parameters[0] == "absolute" {
		path = absolutePath
	}

	return path, err
}
