package plugins

import (
	"os"
	"path"
	"strings"
)

type Virtualenv struct{}

func (plugin Virtualenv) Prompt(parameters []string) (string, error) {
	directory := os.Getenv("VIRTUAL_ENV")

	if len(directory) == 0 {
		return "", nil
	}

	venv := path.Base(directory)
	venv = strings.TrimPrefix(venv, ".")

	prefix, sufix := plugin.complements(parameters)

	return prefix + venv + sufix, nil
}

func (plugin Virtualenv) complements(parameters []string) (string, string) {
	if len(parameters) <= 0 {
		return "", ""
	}

	if len(parameters) == 1 {
		return parameters[0], ""
	}

	return parameters[0], parameters[1]
}
