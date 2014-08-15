package parsers

import "regexp"

type WhiteSpacesParser struct{}

func (parser WhiteSpacesParser) Parse(prompt PromptLine) (PromptLine, []error) {
	regex := regexp.MustCompile(" {2,}")
	prompt.Text = regex.ReplaceAllString(prompt.Text, " ")

	return PromptLine{prompt.Text}, nil
}
