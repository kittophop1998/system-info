package main

import (
	"log"
	"os"
	"system-i/infrastructure/config"
)

func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	cfg, err := config.LoadConfig(env)
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}

	log.Printf("✅ Config loaded successfully from %s.yaml", env)

	initializeApp(cfg)
}
