package utils

import (
	"time"

	"github.com/covertbert/iex-cli/errors"
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
		errors.Error("Failed to convert date string", err)
	}

	return t.Format(time.RFC1123)
}

// ShortDate trims the year from a date in yyyy-mm-dd format and returns
// it in dd/mm format
func ShortDate(d string) string {
	layout := "2006-01-02"
	t, err := time.Parse(layout, d)

	if err != nil {
		errors.Error("Failed to convert date string", err)
	}

	return t.Format("02/01")
}
