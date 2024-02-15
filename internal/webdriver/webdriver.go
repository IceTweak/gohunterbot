package webdriver

import (
	"gohunterbot/internal/logger"
	"os"

	"github.com/tebeka/selenium"
	"go.uber.org/zap"
)

const (
	seleniumPath     = "vendor/selenium-server.jar"
	chromeDriverPath = "vendor/chromedriver.zip"
	Port             = 8080
)

func StartSeleniumService() {
	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),             // Start an X frame buffer for the browser to run in.
		selenium.ChromeDriver(chromeDriverPath), // Specify the path to ChromeDriver in order to use Google.
		selenium.Output(os.Stderr),              // Output debug information to STDERR.
	}
	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, Port, opts...)
	if err != nil {
		logger.Sugar.Errorf("selenium: %w", zap.Error(err))
	}
	defer service.Stop()
}
