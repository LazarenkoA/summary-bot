package main

import (
	"fmt"
	"log/slog"
	"os"
	"summary-bot/bot"
	"summary-bot/storage"
)

var (
	redisaddr = os.Getenv("REDIS")
	botToken  = os.Getenv("BotToken")
	authKey   = os.Getenv("GIGA_AUTH_KEY")
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	if botToken == "" {
		logger.Error("в переменных окружения не задан BotToken")
		os.Exit(1)
	}
	if redisaddr == "" {
		logger.Error("в переменных окружения не задан адрес redis")
		os.Exit(1)
	}

	db, err := storage.NewRedis(redisaddr)
	if err != nil {
		logger.Error(fmt.Sprintf("redis connect error: %v", err))
		os.Exit(1)
	}

	b, err := bot.NewSummaryBot(botToken, authKey, db)
	if err != nil {
		logger.Error(fmt.Sprintf("bot create error: %v", err))
		os.Exit(1)
	}

	b.Run()
}

// systemctl stop summarybot.service && mv -f /var/tmp/summary /home/artem/Summary/ && systemctl start summarybot.service && systemctl status summarybot.service
