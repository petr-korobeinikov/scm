package internal

import "os"

func LookupEnvOrDefault(envName, defaultValue string) (value string, found bool) {
	if value, found := os.LookupEnv(envName); found {
		return value, found
	} else {
		return defaultValue, false
	}
}
