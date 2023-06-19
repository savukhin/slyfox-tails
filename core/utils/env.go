package utils

import (
	"os"
	"strconv"
)

func GetEnvDefault(key string, default_val string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return default_val
	}
	return val
}

func GetEnvDefaultInt(key string, default_val int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return default_val
	}
	val_int, err := strconv.Atoi(val)
	if err != nil {
		return default_val
	}

	return val_int
}
