package plugins

import (
	"strings"
	"time"
)

type DateTime struct{}

func (dateTime DateTime) Prompt(parameter string) (string, error) {
	longFormats := map[string]string{
		"yy": "2006",
		"mm": "03",
		"dd": "02",
		"hh": "15",
		"MM": "04",
		"ss": "05",
	}
	shortFormats := map[string]string{
		"y": "06",
		"m": "3",
		"d": "2",
		"h": "3",
		"M": "4",
		"s": "5",
	}

	layout := parameter
	for key, value := range longFormats {
		layout = strings.Replace(layout, key, value, -1)
	}

	for key, value := range shortFormats {
		layout = strings.Replace(layout, key, value, -1)
	}

	return time.Now().Format(layout), nil
}
