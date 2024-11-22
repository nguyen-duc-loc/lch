package utils

import "fmt"

func padZeroToStart(s string) string {
	if len(s) < 2 {
		s = "0" + s
	}

	return s
}

func FormatTime(t int64) string {
	hours := t / 3600
	t -= hours * 3600
	mins := t / 60
	t -= mins * 60
	secs := t

	hoursInText := padZeroToStart(fmt.Sprintf("%d", hours))
	minsInText := padZeroToStart(fmt.Sprintf("%d", mins))
	secsInText := padZeroToStart(fmt.Sprintf("%d", secs))

	return fmt.Sprintf("%s:%s:%s", hoursInText, minsInText, secsInText)
}
