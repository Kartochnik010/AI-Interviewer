package utils

import "time"

func HumanDate(t time.Time) string {
	// Return empty string if time has the zero value
	if t.IsZero() {
		return ""
	}

	// Convert the time to UTC before formatting it.
	return t.UTC().Format("2006-01-02 15:04:05")
}
