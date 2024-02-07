package bot

import (
	"fmt"
	"gohunterbot/internal/collector"
	"time"

	"github.com/gocolly/colly/v2"
)

type BotOption func(bot *Bot) error

type Bot struct {
	ResourceUrl string
	Collector   *colly.Collector
	Refreshing  time.Duration
}

func defaultBot() *Bot {
	c, _ := collector.NewCollector()
	return &Bot{
		ResourceUrl: "",
		Collector:   c,
		Refreshing:  5 * time.Minute,
	}
}

func NewBot(options ...BotOption) (*Bot, error) {
	b := defaultBot()
	for _, option := range options {
		if err := option(b); err != nil {
			return nil, fmt.Errorf("bot: options: %w", err)
		}
	}
	return b, nil
}

func (b *Bot) Start() {
	// TODO - Updating channel for bot
}
