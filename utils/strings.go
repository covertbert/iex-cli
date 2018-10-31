package utils

// ReplaceEmptyValue replaces empty strings with 'N/A'
func ReplaceEmptyValue(value string) string {
	if len(value) > 0 {
		return value
	}

	return "N/A"
}
