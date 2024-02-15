package bot

import (
	"fmt"
	"gohunterbot/internal/webdriver"
	"os"

	"github.com/tebeka/selenium"
)

func PressUpInSearchBtn(bot *Bot) error {
	// Connect to the WebDriver instance running locally.
	wd, err := selenium.NewRemote(bot.BrowserCaps, fmt.Sprintf("http://localhost:%d/wd/hub", webdriver.Port))
	if err != nil {
		return fmt.Errorf("actions: %w", err)
	}
	defer wd.Quit()
	// Navigate to the target page.
	pageUrl := os.Getenv("RESOURCE_URL")
	if err := wd.Get(pageUrl); err != nil {
		return fmt.Errorf("actions: %w", err)
	}
	// Click needed button.
	btn, err := wd.FindElement(selenium.ByCSSSelector, "button[data-qa='resume-update-button']")
	if err != nil {
		return fmt.Errorf("actions: %w", err)
	}
	if err := btn.Click(); err != nil {
		return fmt.Errorf("actions: %w", err)
	}
	return nil
}
