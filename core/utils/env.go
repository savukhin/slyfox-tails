package utils

import "os"

func GetEnvDefault(key string, default_val string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return default_val
	}
	return val
}
