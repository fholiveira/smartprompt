package main

import "fmt"
import (
	. "github.com/fholiveira/smartprompt/plugins"
	"github.com/fholiveira/smartprompt/plugins/location"
)

type Plugin interface {
	Prompt() (string, error)
}

func main() {
	dict := map[string]Plugin{
		"user":              User{},
		"git":               Git{},
		"location":          location.Default{},
		"location:vimstyle": location.VimStyle{},
	}

	for key, plugin := range dict {
		prompt, err := plugin.Prompt()
		if err == nil {
			fmt.Println(key, ":", prompt)
		}
	}
}
