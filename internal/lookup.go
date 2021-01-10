package internal

import "os"

func LookupEnvOrDefault(envName, defaultValue string) string {
	if value, found := os.LookupEnv(envName); found {
		return value
	} else {
		return defaultValue
	}
}
