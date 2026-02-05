package config

import (
	"os"
	"time"
)

type Config struct {
	Server ServerConfig
	Nova   NovaConfig
	Simbad SimbadConfig
}

type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type NovaConfig struct {
	BaseURL string
	APIKey  string
	Timeout time.Duration
}

type SimbadConfig struct {
	BaseURL string
	Timeout time.Duration
}

func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port:         getEnv("SERVER_PORT", "8080"),
			ReadTimeout:  getDuration("SERVER_READ_TIMEOUT", 30*time.Second),
			WriteTimeout: getDuration("SERVER_WRITE_TIMEOUT", 30*time.Second),
		},
		Nova: NovaConfig{
			BaseURL: getEnv("NOVA_BASE_URL", "https://nova.astrometry.net"),
			APIKey:  os.Getenv("NOVA_API_KEY"),
			Timeout: getDuration("NOVA_TIMEOUT", 30*time.Second),
		},
		Simbad: SimbadConfig{
			BaseURL: getEnv("SIMBAD_BASE_URL", "https://simbad.u-strasbg.fr/simbad/sim-tap/sync"),
			Timeout: getDuration("SIMBAD_TIMEOUT", 10*time.Second),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}

func getDuration(key string, defaultValue time.Duration) time.Duration {
	if v := os.Getenv(key); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			return d
		}
	}
	return defaultValue
}
