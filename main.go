package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"summary-bot/bot"
	"summary-bot/storage"
	"summary-bot/utils"
)

var (
	redisaddr   string
	botToken    string
	gigaAuthKey string
	deepseekKEY string
)

func init() {
	_ = godotenv.Load()
	redisaddr = os.Getenv("REDIS")
	botToken = os.Getenv("BotToken")
	gigaAuthKey = os.Getenv("GIGA_AUTH_KEY")
	deepseekKEY = os.Getenv("DS_API_KEY")
}

func main() {
	logger := utils.NewLogger()

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

	b, err := bot.NewSummaryBot(botToken, deepseekKEY, db)
	if err != nil {
		logger.Error(fmt.Sprintf("bot create error: %v", err))
		os.Exit(1)
	}

	b.Run()
}

// systemctl stop summarybot.service && mv -f /var/tmp/summary /home/artem/Summary/ && systemctl start summarybot.service && systemctl status summarybot.service
