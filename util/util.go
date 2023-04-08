package util

import (
	"os"
)

func GetEnvOrFail(envName string) string {
	var envValue = os.Getenv(envName)
	if envValue == "" {
		panic("env: " + envName + " must be not empty")
	}

	return envValue
}
