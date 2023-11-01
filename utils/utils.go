package utils

import "fmt"

const (
	textBlack = iota + 30
	textRed
	textGreen
	textYellow
	textBlue
	textPurple
	textCyan
	textWhite
)

func Purple(str string) string {
	return textColor(textPurple, str)
}
func Yellow(str string) string {
	return textColor(textYellow, str)
}
func Red(str string) string {
	return textColor(textRed, str)
}

func Green(str string) string {
	return textColor(textGreen, str)
}

func textColor(color int, str string) string {
	return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", color, str)
}
