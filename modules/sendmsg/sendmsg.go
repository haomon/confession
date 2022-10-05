package sendmsg

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testproject/confession/env"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// 分析接收到的訊息
func HandleUpdate(update tgbotapi.Update) {
	text := update.Message.Text
	sticker := update.Message.Sticker

	// userchatID := update.Message.Chat.ID
	// fmt.Println(userchatID)
	// fmt.Println(chatID)
	switch {
	case text != "":
		if strings.HasPrefix(text, env.ChatPrefix) {
			bot.AnalyzeText(update)
		}
	case sticker != nil:
		fmt.Println(sticker.FileID)
		// replyMsg := tgbotapi.NewStickerShare(env.LogchatID, "CAACAgUAAxkBAAOUYzpShhs4JzzEK5cEDqSe0e1KzeYAAmwBAAJvL7ojW2D8vo04eqwqBA")
		// replyMsg.ReplyToMessageID = update.Message.MessageID
		// _, _ = bot.Send(replyMsg)
	}

}

//分析 「ㄡ~修女啊，」 後該做什麼
func AnalyzeText(update tgbotapi.Update) {
	text := update.Message.Text
	text = text[len(env.ChatPrefix):]
	chatID := update.Message.Chat.ID
	if strings.HasPrefix(text, "http") {
		res, err := http.Get("http://i1.hdslb.com/bfs/archive/ec20e0f5fbaa5f8252bf761e4b359edbb767fc43.png")
		if err != nil {
			panic(err)
		}
		photoBytes, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		bot.SendImage(chatID, photoBytes)
	}

	replyMsg := tgbotapi.NewMessage(chatID, "你的意思是說，"+text+"。不過沒關係，一切都好起來的。")

	res, err := http.Get("http://i1.hdslb.com/bfs/archive/ec20e0f5fbaa5f8252bf761e4b359edbb767fc43.png")
	if err != nil {
		panic(err)
	}
	photoBytes, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	// 本地圖片
	// photoBytes, err := ioutil.ReadFile("./img/beargi.png")
	// if err != nil {
	// 	panic(err)
	// }
	photoFileBytes := tgbotapi.FileBytes{
		Name:  "picture",
		Bytes: photoBytes,
	}
	replyImg := tgbotapi.NewPhotoUpload(chatID, photoFileBytes)
	// replyImg.Caption = "sssssss"

	_, _ = bot.Botapi.Send(replyMsg)
	_, _ = bot.Botapi.Send(replyImg)
}

//處理送往訊息管理群的訊息
func LogText(msg string) {
	NewMsg := tgbotapi.NewMessage(env.LogchatID, msg)
	NewMsg.ParseMode = tgbotapi.ModeHTML //傳送html格式的訊息
	_, err := bot.Botapi.Send(NewMsg)
	if err == nil {
		fmt.Println("Send telegram message success")
	} else {
		fmt.Println("Send telegram message error")
		fmt.Println(err)
	}
}

//不用不用報時
func RespNowTime() {
	six := "CAACAgUAAxkBAAOdYzpT_6-vI6-pxFQBkpcPBousZSkAAmkBAAJvL7ojsgABYsiGyO-AKgQ"
	nine := "CAACAgUAAxkBAAOUYzpShhs4JzzEK5cEDqSe0e1KzeYAAmwBAAJvL7ojW2D8vo04eqwqBA"
	twelve := "CAACAgUAAxkBAAObYzpTuZ999NCn4vZdy5BDRzkJAVMAAm8BAAJvL7ojAV9DSiMBSasqBA"
	ten := "CAACAgUAAxkBAAPCYzqH61H_i2x6yTaKck_Lv0PetD4AAm0BAAJvL7ojRdP1rSA7LnwqBA"
	fridayTen := "CAACAgUAAxkBAAOfYzpUZR-HubpnmgABHGePT8ayrnvjAAJVBQACby-6I5UH8HAxbvwCKgQ"
	switch time.Now().Hour() {
	case 9:
		replyMsg := tgbotapi.NewStickerShare(env.ChatID, nine)
		_, err := bot.Botapi.Send(replyMsg)
		if err != nil {
			bot.LogText(err.Error())
		}
	case 12:
		replyMsg := tgbotapi.NewStickerShare(env.ChatID, twelve)
		_, err := bot.Botapi.Send(replyMsg)
		if err != nil {
			bot.LogText(err.Error())
		}
	case 18:
		replyMsg := tgbotapi.NewStickerShare(env.ChatID, six)
		_, err := bot.Botapi.Send(replyMsg)
		if err != nil {
			bot.LogText(err.Error())
		}
	case 22:
		if time.Now().Weekday() == time.Friday {
			replyMsg := tgbotapi.NewStickerShare(env.ChatID, fridayTen)
			_, err := bot.Botapi.Send(replyMsg)
			if err != nil {
				bot.LogText(err.Error())
			}
		} else {
			replyMsg := tgbotapi.NewStickerShare(env.ChatID, ten)
			_, err := bot.Botapi.Send(replyMsg)
			if err != nil {
				bot.LogText(err.Error())
			}
		}
	}
}
