package location

import (
	"os"
	"os/user"
	"strings"
)

type Default struct{}

var homeDirectory = func() (string, error) {
	user, err := user.Current()
	if nil != err {
		return "", err
	}

	return user.HomeDir, nil
}

var currentDirectory = func() (string, string, error) {
	path, err := os.Getwd()
	if nil != err {
		return "", "", err
	}

	home, err := homeDirectory()
	if nil != err {
		return "", "", err
	}

	absolutePath := strings.Replace(path, home, "~", 1)
	symlinkPath := strings.Replace(os.Getenv("PWD"), home, "~", 1)

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
