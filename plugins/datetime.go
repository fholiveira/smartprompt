package plugins

import (
	"strings"
	"time"
)

var now = func() time.Time { return time.Now() }

type DateTime struct{}

func (dateTime DateTime) ShortFormats() map[string]string {
	return map[string]string{
		"y": "06",
		"m": "1",
		"d": "2",
		"h": "03",
		"M": "4",
		"s": "5",
	}
}

func (dateTime DateTime) LongFormats() map[string]string {
	return map[string]string{
		"yy": "2006",
		"mm": "01",
		"dd": "02",
		"hh": "15",
		"MM": "04",
		"ss": "05",
	}
}

func (dateTime DateTime) Prompt(parameters []string) (string, error) {
	layout := parameters[0]

	for key, value := range dateTime.LongFormats() {
		layout = strings.Replace(layout, key, value, -1)
	}

	for key, value := range dateTime.ShortFormats() {
		layout = strings.Replace(layout, key, value, -1)
	}

	return now().Format(layout), nil
}
