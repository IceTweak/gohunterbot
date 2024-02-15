package bot

import (
	"time"

	"github.com/tebeka/selenium"
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

func WithBrowserCaps(caps selenium.Capabilities) BotOption {
	return func(bot *Bot) error {
		bot.BrowserCaps = caps
		return nil
	}
}
