package internal

import "os"

// LookupEnvOrDefault returns value from env or default value if var not set.
func LookupEnvOrDefault(key, defaultValue string) (value string, found bool) {
	if value, found := os.LookupEnv(key); found {
		return value, found
	}

	return defaultValue, false
}
