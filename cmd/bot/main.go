package main

import (
	"context"
	"log"

	"github.com/StezX433/slack-age-bot/internal/config"
	"github.com/StezX433/slack-age-bot/internal/slack"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	bot := slack.NewClient(cfg)

	slack.RegisterCommands(bot)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("Slack Age Bot is starting...")

	if err := bot.Listen(ctx); err != nil {
		log.Fatal(err)
	}
}
