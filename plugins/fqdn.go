package plugins

type FullQualifiedDomainName struct{}

func (fqdn FullQualifiedDomainName) Prompt(parameters []string) (string, error) {
	return "\\H", nil
}
