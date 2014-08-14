package parsers

type Parser interface {
	Parse(prompt string) (string, error)
}
