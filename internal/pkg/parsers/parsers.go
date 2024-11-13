package parsers

import "strconv"

// Helper to parse timeout values with a default
func ParseTimeout(value string, defaultValue int) int {
	timeout, err := strconv.Atoi(value)

	if err != nil {
		return defaultValue
	}

	return timeout
}
