package bot

import (
	"fmt"
	"gohunterbot/internal/logger"
	"gohunterbot/internal/webdriver"
	"time"

	"github.com/tebeka/selenium"
	"go.uber.org/zap"
)

type BotOption func(bot *Bot) error

type Bot struct {
	ResourceUrl string
	BrowserCaps selenium.Capabilities
	Refreshing  time.Duration
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
	webdriver.StartSeleniumService()
	ticker := time.NewTicker(b.Refreshing)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				logger.Sugar.Debug("bot: register new tick")
				err := PressUpInSearchBtn(b)
				if err != nil {
					logger.Sugar.Error("bot", zap.Error(err))
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func defaultBot() *Bot {
	caps := selenium.Capabilities{"browserName": "chrome"}
	return &Bot{
		ResourceUrl: "",
		BrowserCaps: caps,
		Refreshing:  4 * time.Hour,
	}
}
