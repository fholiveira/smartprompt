package plugins

import (
	"os"
	"path"
	"strings"
)

type Virtualenv struct{}

func (plugin Virtualenv) Prompt(parameter string) (string, error) {
	directory := os.Getenv("VIRTUAL_ENV")

	if len(directory) == 0 {
		return "", nil
	}

	venv := path.Base(directory)
	venv = strings.TrimPrefix(venv, ".")

	prefix, sufix := plugin.get_prefix_and_sufix(parameter)

	return prefix + venv + sufix, nil
}

func (plugin Virtualenv) get_prefix_and_sufix(parameter string) (string, string) {
	if len(parameter) <= 0 {
		return "", ""
	}

	if !strings.Contains(parameter, ",") {
		return parameter, ""
	}

	surround := strings.Split(parameter, ",")
	return surround[0], surround[1]
}
