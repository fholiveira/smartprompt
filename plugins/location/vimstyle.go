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

func (plugin VimStyle) applyVimStyle(path string) string {
	if path == "/" || len(path) == 0 {
		return path
	}

	vimPath := ""
	if firstChar(path) == "/" {
		vimPath = "/"
	}

	basePath, directoryName := plugin.splitPath(path)

	for _, directory := range basePath {
		vimPath = gopath.Join(vimPath, firstChar(directory))
	}

	return gopath.Join(vimPath, directoryName)
}

func (plugin VimStyle) Prompt(parameters []string) (string, error) {
	absolutePath, symlinkPath, err := currentDirectory()
	if nil != err {
		return "", err
	}

	path := symlinkPath
	if len(parameters) > 0 && parameters[0] == "absolute" {
		path = absolutePath
	}

	return plugin.applyVimStyle(path), nil
}
