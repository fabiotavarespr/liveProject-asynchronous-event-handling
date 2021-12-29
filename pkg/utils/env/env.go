package env

import (
	"os"
	"strconv"
)

func GetEnvWithDefaultAsString(envKey string, defaultVal string) string {
	val := os.Getenv(envKey)
	if val == "" {
		return defaultVal
	}
	return val
}

func GetEnvWithDefaultAsInt(envKey string, defaultVal int) int {
	val := os.Getenv(envKey)
	if val == "" {
		return defaultVal
	}
	intValue, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}
	return intValue
}
