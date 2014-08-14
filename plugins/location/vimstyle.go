package location

import "strings"

type VimStyle struct{}

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
