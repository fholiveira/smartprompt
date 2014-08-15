package parsers

type ColorParser struct{}

func mapColors() map[string]string {
	return map[string]string{
		"BLACK":             "\\e[0;30m",
		"RED":               "\\e[0;31m",
		"GREEN":             "\\e[0;32m",
		"YELLOW":            "\\e[0;33m",
		"BLUE":              "\\e[0;34m",
		"PURPLE":            "\\e[0;35m",
		"CYAN":              "\\e[0;36m",
		"WHITE":             "\\e[0;37m",
		"BLACK:bold":        "\\e[1;30m",
		"RED:bold":          "\\e[1;31m",
		"GREEN:bold":        "\\e[1;32m",
		"YELLOW:bold":       "\\e[1;33m",
		"BLUE:bold":         "\\e[1;34m",
		"PURPLE:bold":       "\\e[1;35m",
		"CYAN:bold":         "\\e[1;36m",
		"WHITE:bold":        "\\e[1;37m",
		"BLACK:underline":   "\\e[4;30m",
		"RED:underline":     "\\e[4;31m",
		"GREEN:underline":   "\\e[4;32m",
		"YELLOW:underline":  "\\e[4;33m",
		"BLUE:underline":    "\\e[4;34m",
		"PURPLE:underline":  "\\e[4;35m",
		"CYAN:underline":    "\\e[4;36m",
		"WHITE:underline":   "\\e[4;37m",
		"BLACK:background":  "\\e[40m",
		"RED:background":    "\\e[41m",
		"GREEN:background":  "\\e[42m",
		"YELLOW:background": "\\e[43m",
		"BLUE:background":   "\\e[44m",
		"PURPLE:background": "\\e[45m",
		"CYAN:background":   "\\e[46m",
		"WHITE:background":  "\\e[47m",
		"TEXT:reset":        "\\e[0m",
	}
}

func (parser ColorParser) Parse(prompt PromptLine) (PromptLine, []error) {
	colors := mapColors()

	for _, token := range prompt.Tokens() {
		color, isColor := colors[token.Name()]
		if isColor {
			prompt.Apply(token, color)
		}
	}

	return PromptLine{prompt.Text}, nil
}
