// Package config 配置包
package config

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

// GetEnv GetEnv
func GetEnv(key string) (value string) {
	return os.Getenv(key)
}

// GetEnvForError GetEnvForError
func GetEnvForError(key string) (value string, err error) {
	value = os.Getenv(key)
	if len(value) == 0 {
		err = errors.New("Configuration not obtained")
	}
	return
}

// GetEnvForPanic GetEnvForPanic
func GetEnvForPanic(key string) (value string) {
	value = os.Getenv(key)
	if len(value) == 0 {
		panic("Configuration not obtained")
	}
	return
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
	case "TRUE", "true", "T", "t":
		return true
	case "FALSE", "false", "F", "f":
		return false
	default:
		return defaultValue
	}
}
