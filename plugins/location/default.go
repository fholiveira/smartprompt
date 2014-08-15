package location

import (
	"os"
	"os/user"
	"strings"
)

type Default struct{}

func getWorkingDir() (string, error) {
	workingDirectory, err := os.Getwd()
	if nil != err {
		return "", err
	}

	user, err := user.Current()
	if nil != err {
		return "", err
	}

	return strings.Replace(workingDirectory, user.HomeDir, "~", 1), nil
}

func (location Default) Prompt(parameter string) (string, error) {
	return getWorkingDir()
}
