package plugins

import (
	"os"
	"path"
	"strings"
)

var getVirtualenv = func() string { return os.Getenv("VIRTUAL_ENV") }

type Virtualenv struct{}

func (plugin Virtualenv) Prompt(parameters []string) (string, error) {
	directory := getVirtualenv()

	if len(directory) == 0 {
		return "", nil
	}

	prefix, sufix := plugin.complements(parameters)
	venv := strings.TrimPrefix(path.Base(directory), ".")

	return prefix + venv + sufix, nil
}

func (plugin Virtualenv) complements(parameters []string) (string, string) {
	if len(parameters) == 0 {
		return "", ""
	}

	if len(parameters) == 1 {
		return parameters[0], ""
	}

	return parameters[0], parameters[1]
}
