package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SlackBotToken string
	SlackAppToken string
}

func Load() (Config, error) {
	_ = godotenv.Load()

	config := Config{
		SlackBotToken: os.Getenv("SLACK_BOT_TOKEN"),
		SlackAppToken: os.Getenv("SLACK_APP_TOKEN"),
	}

	if config.SlackBotToken == "" {
		return Config{}, fmt.Errorf("SLACK_BOT_TOKEN is not set")
	}

	if config.SlackAppToken == "" {
		return Config{}, fmt.Errorf("SLACK_APP_TOKEN is not set")
	}

	return config, nil
}
