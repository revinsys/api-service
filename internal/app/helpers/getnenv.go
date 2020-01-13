package helpers

import "os"

//GetEnv ...
func GetEnv(key, fallback string) string {
	if key, ok := os.LookupEnv(key); ok {
		return key
	}

	return fallback
}