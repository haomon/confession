package tgbot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	Botapi *tgbotapi.BotAPI
}

func New(tgToken string) *Bot {
	var err error
	bot := &Bot{
		Botapi: nil,
	}
	bot.Botapi, err = tgbotapi.NewBotAPI(tgToken)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	bot.Botapi.Debug = false
	return bot
}

func (bot *Bot) SendText(chatID int64, text string) {
	replyMsg := tgbotapi.NewMessage(chatID, text)
	// replyImg.Caption = "sssssss"
	_, _ = bot.Botapi.Send(replyMsg)
}

func (bot *Bot) SendImage(chatID int64, dataByte []byte) {
	photoFileBytes := tgbotapi.FileBytes{
		Name:  "picture",
		Bytes: dataByte,
	}
	replyImg := tgbotapi.NewPhotoUpload(chatID, photoFileBytes)
	_, _ = bot.Botapi.Send(replyImg)
}

func (bot *Bot) SendImageAddMsg(chatID int64, dataByte []byte, msg string) {
	photoFileBytes := tgbotapi.FileBytes{
		Name:  "picture",
		Bytes: dataByte,
	}
	replyImg := tgbotapi.NewPhotoUpload(chatID, photoFileBytes)
	replyImg.Caption = msg
	_, _ = bot.Botapi.Send(replyImg)
}
