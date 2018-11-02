package utils

import (
	"fmt"
	"time"
)

// UNIXToHumanReadable converts a UNIX timestamp to human-readable 24 hour time
func UNIXToHumanReadable(u int64) string {
	return time.Unix(u, 0).Format("15:04")
}

// DateStringToHumanReadable converts date string to human-readable day and time
func DateStringToHumanReadable(d string) string {
	layout := "2006-01-02T15:04:05-07:00"
	t, err := time.Parse(layout, d)

	if err != nil {
		fmt.Println(err)
	}

	return t.Format(time.RFC1123)
}
