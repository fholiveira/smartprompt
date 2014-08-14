package parsers

import "regexp"

type WhiteSpacesParser struct{}

func (parser WhiteSpacesParser) Parse(prompt string) (string, error) {
	regex := regexp.MustCompile(" {2,}")
	return regex.ReplaceAllString(prompt, " "), nil
}
