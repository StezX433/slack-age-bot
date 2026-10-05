package slack

import (
	"github.com/shomali11/slacker"

	"github.com/StezX433/slack-age-bot/internal/config"
)

func NewClient(cfg config.Config) *slacker.Slacker {
	return slacker.NewClient(
		cfg.SlackBotToken,
		cfg.SlackAppToken,
	)
}
