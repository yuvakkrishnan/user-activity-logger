// Configuration loader
package config

import "os"

// TODO: Load from env or config file

func GetEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
