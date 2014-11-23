package location

import (
	gopath "path"
	"strings"
)

func firstChar(text string) string {
	return string([]rune(text)[0])
}

type VimStyle struct{}

func (plugin VimStyle) splitPath(directory string) ([]string, string) {
	path := []string{}
	for _, dir := range strings.Split(directory, "/") {
		if dir != "" {
			path = append(path, dir)
		}
	}
	lastItem := len(path) - 1
	return path[:lastItem], path[lastItem]
}

func (plugin VimStyle) concatWithVimStyle(workingDirectory string) string {
	if workingDirectory == "/" {
		return workingDirectory
	}

	path, directory := plugin.splitPath(workingDirectory)

	var vimPath string
	for _, directory := range path {
		vimPath = gopath.Join(vimPath, firstChar(directory))
	}

	if firstChar(workingDirectory) == "/" {
		vimPath = "/" + vimPath
	}

	return gopath.Join(vimPath, directory)
}

func (plugin VimStyle) Prompt(parameters []string) (string, error) {
	workingDirectory, err := getWorkingDir()
	if nil != err {
		return "", err
	}

	return plugin.concatWithVimStyle(workingDirectory), nil
}
