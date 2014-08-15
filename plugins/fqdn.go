package plugins

type FullQualifiedDomainName struct{}

func (fqdn FullQualifiedDomainName) Prompt() (string, error) {
	return "\\H", nil
}
