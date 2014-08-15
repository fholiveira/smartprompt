package parsers

type Parser interface {
	Parse(prompt PromptLine) (PromptLine, error)
}
