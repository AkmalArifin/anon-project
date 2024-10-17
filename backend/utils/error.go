package utils

import "regexp"

// Duplicate entry 'ikanikan' for key 'users.username' [message, value, field]
func GetDuplicateError(message string) []string {
	pattern := `1062.*'(.+)'.*'([\w.]+)'`
	re := regexp.MustCompile(pattern)

	match := re.FindStringSubmatch(message)
	return match
}
