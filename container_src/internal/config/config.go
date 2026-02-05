package config

import (
	"os"
	"time"
)

type Config struct {
	Server ServerConfig
	Nova   NovaConfig
	Simbad SimbadConfig
	KV     KVConfig
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

type KVConfig struct {
	BaseURL     string
	AccountID   string
	NamespaceID string
	APIToken    string
	Enabled     bool
	Timeout     time.Duration
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
		KV: loadKVConfig(),
	}
}

func loadKVConfig() KVConfig {
	accountID := os.Getenv("CF_ACCOUNT_ID")
	namespaceID := os.Getenv("CF_KV_NAMESPACE_ID")
	apiToken := os.Getenv("CF_API_TOKEN")

	return KVConfig{
		BaseURL:     getEnv("CF_KV_BASE_URL", "https://api.cloudflare.com/client/v4"),
		AccountID:   accountID,
		NamespaceID: namespaceID,
		APIToken:    apiToken,
		Enabled:     accountID != "" && namespaceID != "" && apiToken != "",
		Timeout:     getDuration("KV_TIMEOUT", 5*time.Second),
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
