package config

import (
	"os"
	"strconv"
	"strings"
)

// GetEnv GetEnv
func GetEnv(key string) (value string) {
	return os.Getenv(key)
}

// GetDefaultEnv GetDefaultEnv
func GetDefaultEnv(key, defaultValue string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// GetEnvToInt GetEnvToInt
func GetEnvToInt(key string) (value int) {
	tmp := os.Getenv(key)
	if tmp == "" {
		return 0
	}
	value, _ = strconv.Atoi(tmp)
	return value
}

// GetDefaultEnvToInt GetDefaultEnvToInt
func GetDefaultEnvToInt(key string, defaultValue int) (value int) {
	tmp := os.Getenv(key)
	if tmp == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(tmp)
	if err != nil {
		return defaultValue
	}
	return value
}

// GetEnvToBool GetEnvToBool
func GetEnvToBool(key string) (value bool) {
	switch strings.ToUpper(os.Getenv(key)) {
	case "TRUE":
		return true
	default:
		return false
	}
}

// GetDefaultEnvToBool GetDefaultEnvToBool
func GetDefaultEnvToBool(key string, defaultValue bool) (value bool) {
	tmp := os.Getenv(key)
	if tmp == "" {
		return defaultValue
	}
	switch strings.ToUpper(tmp) {
	case "TRUE":
		return true
	case "FALSE":
		return false
	default:
		return defaultValue
	}
}
