package parsers

import "github.com/fholiveira/smartprompt/colors"

type ColorParser struct{}

func (parser ColorParser) Colors() map[string]colors.Color {
	return map[string]colors.Color{
		"YELLOW": colors.Yellow,
		"PURPLE": colors.Purple,
		"BLACK":  colors.Black,
		"GREEN":  colors.Green,
		"WHITE":  colors.White,
		"BLUE":   colors.Blue,
		"CYAN":   colors.Cyan,
		"RED":    colors.Red,
	}
}

func (parser ColorParser) apply(color colors.Color, token Token) string {
	parameters := token.Parameters()
	if len(parameters) > 0 {
		switch parameters[0] {
		case "bold":
			return color.Bold()
		case "underline":
			return color.Underline()
		case "background":
			return color.Background()
		}
	}

	return color.Normal()
}

func (parser ColorParser) Parse(prompt PromptLine) (PromptLine, []error) {
	colorsMap := parser.Colors()

	for _, token := range prompt.Tokens(":") {
		color, isColor := colorsMap[token.Name()]
		if isColor {
			prompt.Apply(token, parser.apply(color, token))
		}

		if token.Name() == "TEXT" {
			prompt.Apply(token, colors.TextReset)
		}
	}

	return PromptLine{prompt.Text}, nil
}
