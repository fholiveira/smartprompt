package colors

import "strconv"

type Color int

func (color Color) Bold() string {
	return color.toString("1;")
}

func (color Color) Normal() string {
	return color.toString("0;")
}

func (color Color) Underline() string {
	return color.toString("4;")
}

func (color Color) Background() string {
	return color.toString("")
}

func (color Color) toString(formatting string) string {
	return "\\[\\e[" + formatting + strconv.Itoa(int(color)) + "m\\]"
}
