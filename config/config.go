package config

import "os"

type Config struct {
	Key string
}

func GetConfig() *Config {
	res := os.Getenv("API_KEY")
	if res == "" {
		panic("No API_KEY in .env")
	}
	return &Config{Key: res}
}
