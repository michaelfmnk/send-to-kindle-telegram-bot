package main

import (
	"github.com/michaelfmnk/send-to-kindle-telegram-bot/bot"
	"log"
	"os"
)

func main() {
	unkindleBot := bot.SendToKindleBot{
		Token:     os.Getenv("UBOT_TELEGRAM_TOKEN"),
		EmailFrom: os.Getenv("UBOT_EMAIL_FROM"),
		EmailTo:   os.Getenv("UBOT_EMAIL_TO"),
		SmtpHost:  os.Getenv("UBOT_SMTP_HOST"),
		Password:  os.Getenv("UBOT_PASSWORD"),
	}
	if err := unkindleBot.Start(); err != nil {
		log.Fatal("could not start telegram bot", err)
	}
}
