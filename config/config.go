package config

import "os"

type Config struct {
	Key string
}

func GetConfig() *Config {
	res := os.Getenv("KEY")
	if res == "" {
		panic("No KEY in .env")
	}
	return &Config{Key: res}
}
