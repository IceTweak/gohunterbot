package main

import (
	"context"
	"gohunterbot/internal/bot"
	"gohunterbot/internal/logger"
	"os"

	"go.uber.org/zap"
	"google.golang.org/appengine/log"
)

func main() {
	ctx := context.Background()
	err := logger.InitSugar()
	if err != nil {
		log.Errorf(ctx, "Failed to setup logger: %s", err.Error())
	}
	logger.Sugar.Info("Starting HeadHunter Bot...")

	bot, err := bot.NewBot(bot.WithResUrl(os.Getenv("RESOURCE_URL")))
	if err != nil {
		logger.Sugar.Error(zap.Error(err))
	}
	bot.Start()
}
