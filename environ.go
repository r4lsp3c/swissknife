package swissknife

import "os"

// GetEnvOrDefault
// return environ specified by key if it's set, _default otherwise
func GetEnvOrDefault(key, _default string) (value string) {
	if value = os.Getenv(key); value == "" {
		return _default
	}
	return value
}
