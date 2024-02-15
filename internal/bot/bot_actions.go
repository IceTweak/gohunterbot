package bot

import (
	"fmt"
	"os"

	"github.com/tebeka/selenium"
)

func PressUpInSearchBtn(bot *Bot) error {
	// Connect to the WebDriver instance running in docker container.
	seleniumPort := os.Getenv("SELENIUM_DRIVER_PORT")
	wd, err := selenium.NewRemote(bot.BrowserCaps, fmt.Sprintf("http://localhost:%s/wd/hub", seleniumPort))
	if err != nil {
		return fmt.Errorf("bot: %w", err)
	}
	defer wd.Quit()
	// Navigate to the target page.
	if err := wd.Get(bot.ResourceUrl); err != nil {
		return fmt.Errorf("bot: %w", err)
	}
	// Click needed button.
	btn, err := wd.FindElement(selenium.ByCSSSelector, "button[data-qa='resume-update-button']")
	if err != nil {
		return fmt.Errorf("bot: %w", err)
	}
	if err := btn.Click(); err != nil {
		return fmt.Errorf("bot: %w", err)
	}
	return nil
}
