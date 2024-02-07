package main

import (
	"context"
	"gohunterbot/internal/bot"
	"gohunterbot/internal/logger"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"google.golang.org/appengine/log"
)

//? Implementation Details
//* Bot should visit HeadHunter web-page every Refreshing interval;

//* Bot should have a gocolly Collector for parsing this page
//* and make some activities (press button, fill form);

func main() {
	ctx := context.Background()
	err := logger.InitSugar()
	if err != nil {
		log.Errorf(ctx, "Failed to setup logger: %s", err.Error())
	}
	err = godotenv.Load()
	if err != nil {
		logger.Sugar.Error("dotenv", zap.Error(err))
	}
	logger.Sugar.Info("Starting HeadHunter Bot...")

	bot, err := bot.NewBot(bot.WithResUrl(os.Getenv("RESOURCE_URL")))
	if err != nil {
		logger.Sugar.Error(zap.Error(err))
	}
	bot.Start()
}
