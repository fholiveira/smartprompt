package plugins

type FullQualifiedDomainName struct{}

func (fqdn FullQualifiedDomainName) Prompt(parameter string) (string, error) {
	return "\\H", nil
}
