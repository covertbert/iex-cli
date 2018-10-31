package utils

import (
	"time"
)

// UNIXToHumanReadable converts a UNIX timestamp to human-readable 24 hour time
func UNIXToHumanReadable(u int64) string {
	return time.Unix(u, 0).Format("15:04")
}
