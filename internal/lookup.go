package internal

import "os"

func LookupEnvOrDefault(key, defaultValue string) (value string, found bool) {
	if value, found := os.LookupEnv(key); found {
		return value, found
	} else {
		return defaultValue, false
	}
}
