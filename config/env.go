package config

import "os"

func GetEnv(env string) string {
	getEnv := os.Getenv(env)
	return getEnv
}
