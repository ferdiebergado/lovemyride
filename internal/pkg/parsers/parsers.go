package parsers

import "strconv"

// Helper to parse integer values with a default
func ParseInt(value string, defaultValue int) int {
	timeout, err := strconv.Atoi(value)

	if err != nil {
		return defaultValue
	}

	return timeout
}
