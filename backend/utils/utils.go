package utils

import "os"

// GetEnv is a function that returns the value of an environment variable
func GetEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}

	return defaultVal
}

func IsProd() bool {
	return GetEnv("ENV", "") == "prod"
}

func IsDev() bool {
	return GetEnv("ENV", "") == "dev"
}
