package tg

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/likexian/whois"
)

func Init(TGBotToken string) {

	bot, err := tgbotapi.NewBotAPI(TGBotToken)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Bot Start!")
	log.Println("Bot Name:", bot.Self.UserName)
	log.Println("Bot ID:", bot.Self.ID)

	msg := tgbotapi.NewMessage(735381053, "Bot Start!")
	bot.Send(msg)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		log.Println(update.Message.Chat.ID, update.Message.Text)

		// 勝勝 @gnehs_owo 215616188
		if update.Message.From.ID == 215616188 {
			if update.Message.Text == "勝勝可愛" {
				// 跟小易說
				msg := tgbotapi.NewMessage(735381053, "勝勝剛剛跟我說他可愛")
				bot.Send(msg)

				// 跟勝勝說
				msg = tgbotapi.NewMessage(215616188, "真的真的，勝勝超可愛的！")
				bot.Send(msg)
			} else {
				// 跟小易說
				msgContent := fmt.Sprintf("勝勝剛剛跟我說：\n\n%s", update.Message.Text)
				msg := tgbotapi.NewMessage(735381053, msgContent)
				bot.Send(msg)

				// 跟勝勝說
				msg = tgbotapi.NewMessage(215616188, "Emm... 好吧，我現在不知道這個意思。我會請小易回覆你的。")
				bot.Send(msg)
			}
		}

		// 小易用
		if update.Message.From.ID == 735381053 {
			if update.Message.Command() == "sendToGnehs" {
				// 跟勝勝說
				msgContent := fmt.Sprintf("小易說\n\n%s", update.Message.Text)
				msg := tgbotapi.NewMessage(215616188, msgContent)
				bot.Send(msg)

				// 跟小易說
				msgContent = fmt.Sprintf("我已經把你的訊息傳給勝勝了。\n\n%s", update.Message.Text)
				msg = tgbotapi.NewMessage(735381053, msgContent)
				bot.Send(msg)
			}

			if update.Message.Command() == "sendToYCwouldliketoTainan" {
				// To yc 要揪台南
				msgContent := fmt.Sprintf("小易說\n\n%s", update.Message.Text)
				msg := tgbotapi.NewMessage(-1001655145767, msgContent)
				bot.Send(msg)

				// 跟小易說
				msgContent = fmt.Sprintf("我已經把你的訊息傳到群組了。\n\n%s", update.Message.Text)
				msg = tgbotapi.NewMessage(735381053, msgContent)
				bot.Send(msg)
			}
		}

		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		sendMessageCheck := false
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		switch update.Message.Command() {
		case "whois":
			content := strings.ReplaceAll(update.Message.Text, "/whois", "")
			result, err := whois.Whois(content)
			if err != nil {
				fmt.Println(result)
			}
			msg.Text = fmt.Sprintf("`%s`", result)
		case "mtr":
			// https://github.com/steveyiyo/LookingGlassBot
		case "ping":
			// https://github.com/steveyiyo/LookingGlassBot
		case "traceroute":
			// https://github.com/steveyiyo/LookingGlassBot
		case "pong":
			msg.Text = "pong"
		case "getid":
			msg.Text = fmt.Sprintf("`%d`", update.Message.Chat.ID)
		case "help":
			msg.Text = "使用 /whois 來查詢 WHOIS 訊息。\n使用 /mtr 來查詢 MTR 訊息。\n使用 /ping 來查詢 Ping 訊息。\n使用 /traceroute 來查詢 Traceroute 訊息。"
		default:
			sendMessageCheck = false
			if update.Message.Chat.ID != 735381053 {
				msg.Text = "使用 /help 來查看可使用的指令列表。"
			}
		}
		msg.ParseMode = "markdown"

		if sendMessageCheck {
			if _, err := bot.Send(msg); err != nil {
				fmt.Println(err)
			}
		}
	}
}
