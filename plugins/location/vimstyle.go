package location

import (
	"os"
	"os/user"
	"strings"
)

type VimStyle struct{}

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

func splitDirectory(directory string) []string {
	path := []string{}
	for _, dir := range strings.Split(directory, "/") {
		if dir != "" {
			path = append(path, dir)
		}
	}

	return path
}

func concatWithVimStyle(workingDirectory string) string {
	if workingDirectory == "/" {
		return workingDirectory
	}

	path := splitDirectory(workingDirectory)

	var simplePath string
	for _, directory := range path[:len(path)-1] {
		simplePath += string([]rune(directory)[0]) + "/"
	}

	if string([]rune(workingDirectory)[0]) == "/" {
		simplePath = "/" + simplePath
	}

	simplePath += path[len(path)-1]

	return simplePath
}

func (location VimStyle) Prompt() (string, error) {
	workingDirectory, err := getWorkingDir()
	if nil != err {
		return "", err
	}

	return concatWithVimStyle(workingDirectory), nil
}
