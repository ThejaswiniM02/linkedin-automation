package config

import (
	"os"
)

type Config struct {
	UserAgent string
	Headless  bool
}

func Load() Config {
	return Config{
		UserAgent: getEnv("USER_AGENT", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 Chrome/120.0.0.0 Safari/537.36"),
		Headless:  getEnv("HEADLESS", "false") == "true",
	}
}

func getEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}
