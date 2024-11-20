package utils

import (
	"github.com/fatih/color"
)

func BoldText(s string) string {
	bold := color.New(color.Bold)

	return bold.Sprintf("%s", s)
}
