package location

import (
	"fmt"
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
	fmt.Println(user)
	user, err := user.Current()
	if nil != err {
		return "", err
	}

	return strings.Replace(workingDirectory, user.HomeDir, "~", 1), nil
}

func (location Default) Prompt() (string, error) {
	return getWorkingDir()
}
