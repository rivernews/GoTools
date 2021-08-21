package main

import (
	"os"
	"strings"
)

// getEnvVarHelper - don't care about no value when getting env var.
// Do not use this for credential, because we should always make sure credentials are available
// to avoid comparing to empty string when auth
func GetEnvVarHelper(key string) string {
	return getEnvVarOrDefault(key, "")
}
func getBoolEnvVarHelper(key string) bool {
	value := strings.TrimSpace(strings.ToLower(GetEnvVarHelper(key)))
	if value == "true" || value == "yes" || value == "1" {
		return true
	}
	return false
}

// getEnvVarOrDefault - must give a default value
func getEnvVarOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return defaultValue
}