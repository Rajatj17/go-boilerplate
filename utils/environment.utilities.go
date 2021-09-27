package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type EnvVariableUtil struct{}

func (e EnvVariableUtil) GetFieldFromEnv(args ...string) string {
	print(args)
	if len(args) == 0 {
		return ""
	}

	fieldName := args[0]

	value, exists := os.LookupEnv(fieldName)
	if !exists {
		if len(args) < 1 {
			return ""
		}

		value = args[1]
	}

	return value
}

// Method for Sanitizing String
func (e EnvVariableUtil) GetSanitizedEnvAsString(args ...string) string {
	print(args)
	value := e.GetFieldFromEnv(args...)

	trimmedValue := strings.TrimSpace(value)

	return trimmedValue
}

// Method for Santitizing Ints
func (e EnvVariableUtil) GetSanitizedEnvAsInt(args ...string) int {
	value := e.GetFieldFromEnv(args...)

	intVal, conversionError := strconv.Atoi(value)
	if conversionError != nil {
		panicMessage := fmt.Sprintf("Error: Env value %s can't be converted to integer", value)
		panic(panicMessage)
	}

	return intVal
}

// Method for Sanitizing Boolean
func (e EnvVariableUtil) GetSanitizedEnvAsBoolean(args ...string) bool {
	value := e.GetFieldFromEnv(args...)

	boolVal, conversionError := strconv.ParseBool(value)
	if conversionError != nil {
		return false
	}

	return boolVal
}
