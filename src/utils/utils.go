package utils

import "os"

func GetEnvVariable(key string, def string) string {
	value := os.Getenv(key)
	if value == "" {
		return def
	}
	return value
}
