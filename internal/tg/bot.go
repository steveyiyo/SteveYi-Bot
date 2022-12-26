package tg

import (
	"fmt"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/likexian/whois"
)

func Init() {
	godotenv.Load()

	TGBotToken := os.Getenv("TelegramBotToken")

	bot, err := tgbotapi.NewBotAPI(TGBotToken)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success!")
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		log.Println(update.Message.Chat.ID, update.Message.Text)
		switch update.Message.Command() {
		case "whois":
			content := strings.ReplaceAll(update.Message.Text, "/whois", "")
			result, err := whois.Whois(content)
			if err != nil {
				fmt.Println(result)
			}
			msg.Text = "`" + result + "`"
		case "mtr":
			// https://github.com/steveyiyo/LookingGlassBot
		case "ping":
			// https://github.com/steveyiyo/LookingGlassBot
		case "traceroute":
			// https://github.com/steveyiyo/LookingGlassBot
		default:
			msg.Text = "使用 /help 來查看可使用的指令列表。"
		}
		msg.ParseMode = "markdown"

		if _, err := bot.Send(msg); err != nil {
			fmt.Println(err)
		}
	}
}
