package utils

import (
	"github.com/fatih/color"
)

func BoldText(s string) string {
	bold := color.New(color.Bold)

	return bold.Sprintf("%s", s)
}

func RedText(s string) string {
	red := color.New(color.FgRed)

	return red.Sprintf("%s", s)
}

func OrangeText(s string) string {
	orange := color.RGB(255, 128, 0)

	return orange.Sprintf("%s", s)
}

func BlueText(s string) string {
	blue := color.New(color.FgBlue)

	return blue.Sprintf("%s", s)
}

func MagentaText(s string) string {
	magenta := color.New(color.FgMagenta)

	return magenta.Sprintf("%s", s)
}

func GreenText(s string) string {
	green := color.New(color.FgGreen)

	return green.Sprintf("%s", s)
}

func WhiteText(s string) string {
	white := color.New(color.FgWhite)

	return white.Sprintf("%s", s)
}
