package location

import (
	"os"
	gopath "path"
	"strings"
)

var homeDir, _ = homeDirectory()

func firstChar(text string) string {
	return string([]rune(text)[0])
}

func isSymlink(path string) bool {
	fileInfo, err := os.Lstat(strings.Replace(path, "~", homeDir, 1))
	if nil != err {
		return false
	}

	return fileInfo.Mode()&os.ModeSymlink > 0
}

type VimStyle struct{}

func (plugin VimStyle) splitPath(path string) ([]string, string) {
	path, directory := gopath.Split(path)
	newPath := []string{}
	for _, dir := range strings.Split(path, "/") {
		if dir != "" {
			newPath = append(newPath, dir)
		}
	}

	return newPath, directory
}

func (plugin VimStyle) applyStyle(path string, symlinks bool) string {
	if path == "/" || len(path) == 0 {
		return path
	}

	normalPath, vimPath := "", ""
	if firstChar(path) == "/" {
		normalPath, vimPath = "/", "/"
	}

	basePath, dirName := plugin.splitPath(path)
	for _, directory := range basePath {
		normalPath = gopath.Join(normalPath, directory)

		alias := firstChar(directory)
		if symlinks && isSymlink(normalPath) {
			alias += "@"
		}

		vimPath = gopath.Join(vimPath, alias)
	}

	normalPath = gopath.Join(normalPath, dirName)
	if symlinks && isSymlink(normalPath) {
		dirName += "@"
	}

	return gopath.Join(vimPath, dirName)
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

	showSymlinks := true
	if len(parameters) > 1 && parameters[1] == "false" {
		showSymlinks = false
	}

	return plugin.applyStyle(path, showSymlinks), nil
}
