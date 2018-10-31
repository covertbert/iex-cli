package utils

// ReplaceEmpty replaces empty strings with 'N/A'
func ReplaceEmpty(value string) string {
	if len(value) > 0 {
		return value
	}

	return "N/A"
}
