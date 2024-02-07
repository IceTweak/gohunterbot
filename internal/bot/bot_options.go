package bot

import (
	"time"

	"github.com/gocolly/colly/v2"
)

func WithResUrl(url string) BotOption {
	return func(bot *Bot) error {
		bot.ResourceUrl = url
		return nil
	}
}

func WithRefreshingInterval(interval time.Duration) BotOption {
	return func(bot *Bot) error {
		bot.Refreshing = interval
		return nil
	}
}

func WithCustomCollector(collector *colly.Collector) BotOption {
	return func(bot *Bot) error {
		bot.Collector = collector
		return nil
	}
}
