package utils

import (
	"fmt"
	"time"
)

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

func suffixS(key string, value int64) string {
	output := fmt.Sprintf("%d %s", value, key)
	if value > 1 {
		output += "s"
	}
	return output
}

func FormatSince(t int64) string {
	convertedTime := time.Unix(t, 0)
	diff := int64(time.Since(convertedTime).Seconds())

	output := ""
	days := diff / 86400
	hours := diff / 3600
	mins := diff / 60

	if days > 0 {
		output += suffixS("day", days)
	} else if hours > 0 {
		output += suffixS("hour", hours)
	} else if mins > 0 {
		output += suffixS("minute", mins)
	} else {
		output += "A few seconds"
	}

	if diff > 0 {
		output += " ago"
	}

	return output
}
