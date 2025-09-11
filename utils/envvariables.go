package utils

import (
	"fmt"
	"os"
	"strings"
)

func GetEnvironmentVariables(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" || len(strings.TrimSpace(value)) == 0 {
		return "", fmt.Errorf("missing required environment variable: %s", key)
	}
	return value, nil
}
