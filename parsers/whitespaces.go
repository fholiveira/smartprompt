package parsers

import "regexp"

type WhiteSpacesParser struct{}

const whiteSpace = " "

func (parser WhiteSpacesParser) Parse(prompt PromptLine) (PromptLine, []error) {
	regex := regexp.MustCompile(" {2,}")
	text := regex.ReplaceAllString(prompt.Text, whiteSpace)

	return PromptLine{text}, nil
}
