package env

import (
	"os"
	"strconv"
)

func GetString(key, placeholder string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return placeholder
	}
	return val
}

func GetInt(key string, placeholder int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return placeholder
	}
	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return placeholder
	}

	return valAsInt
}

func GetBool(key string, placeholder bool) bool {
	val, ok := os.LookupEnv(key)
	if !ok {
		return placeholder
	}
	boolVal, err := strconv.ParseBool(val)
	if err != nil {
		return placeholder
	}

	return boolVal
}
