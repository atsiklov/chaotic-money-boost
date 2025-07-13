package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func loadEnvToMemory() {
	// dir, _ := os.Getwd()
	// log.Println("👍Loading .env. Current working directory:", dir)

	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env file not found, continuing without it")
	}
}

func overrideWithEnv(cfg *AppConfig) {
	envMap := []struct {
		key   string
		value *string
	}{
		{"DB_HOST", &cfg.DbServer.Host},
		{"DB_PORT", &cfg.DbServer.Port},
		{"DB_USER", &cfg.DbServer.User},
		// {"DB_PASS", &cfg.DbServer.Pass},
		{"DB_NAME", &cfg.DbServer.Name},
	}

	var missing []string

	for _, cfgParam := range envMap {
		envKey := cfgParam.key
		envVal := os.Getenv(envKey)

		if envVal == "" {
			missing = append(missing, envKey)
		} else {
			*cfgParam.value = envVal
		}
	}

	if len(missing) > 0 {
		log.Fatalf("❌ Missing required environment variables: %v", strings.Join(missing, ", "))
	}
}
