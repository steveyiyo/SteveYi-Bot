package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/steveyiyo/steveyi-bot/internal/tg"
)

func main() {
	godotenv.Load()

	TGBotToken := os.Getenv("TelegramBotToken")

	if TGBotToken == "" {
		log.Fatalln("Telegram Bot Token is not set")
	}

	tg.Init(TGBotToken)
}
