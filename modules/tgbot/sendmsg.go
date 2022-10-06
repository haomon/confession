package tgbot

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"testproject/confession/env"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// 分析接收到的訊息
func imageByte(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	bytes, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	return bytes
}

// 分析接收到的訊息
func (bot *Bot) HandleUpdate(update tgbotapi.Update) {
	text := update.Message.Text
	sticker := update.Message.Sticker
	photo := update.Message.Photo

	// userchatID := update.Message.Chat.ID
	// fmt.Println(userchatID)
	// fmt.Println(chatID)
	switch {
	case text != "":
		if strings.HasPrefix(text, env.ChatPrefix) {
			bot.AnalyzeText(update)
		}
	case sticker != nil:
		// fmt.Println(sticker.FileID)
		// replyMsg := tgbotapi.NewStickerShare(env.LogchatID, "CAACAgUAAxkBAAOUYzpShhs4JzzEK5cEDqSe0e1KzeYAAmwBAAJvL7ojW2D8vo04eqwqBA")
		// replyMsg.ReplyToMessageID = update.Message.MessageID
		// _, _ = bot.Send(replyMsg)
	case photo != nil:
		fmt.Println(update.Message.Caption)
		caption := update.Message.Caption
		if len(caption) > 0 && strings.HasPrefix(caption, "安妮亞喜歡這個修女") {
			bot.killlakillnun(update)
		}
	}

}

//分析 「ㄡ~修女啊，」 後該做什麼
func (bot *Bot) AnalyzeText(update tgbotapi.Update) {
	text := update.Message.Text
	text = text[len(env.ChatPrefix):]
	chatID := update.Message.Chat.ID

	var lotton bool
	if strings.Contains(text, "[") && strings.Contains(text, "]") {
		lotton = true
	}

	var respmsg string
	switch {
	case strings.HasPrefix(text, "http"):
		photoBytes := imageByte("https://memeprod.sgp1.digitaloceanspaces.com/user-template/4cf02807f8248a46d6f0ded47595a923.png")
		bot.SendImage(chatID, photoBytes)
		return
	case strings.HasPrefix(text, "救救我") && lotton:
		li := strings.Index(text, "[")
		ri := strings.Index(text, "]")
		lottoString := text[li+1 : ri]
		lotto := strings.Split(lottoString, " ")

		result := lotto[rand.Intn(len(lotto))]
		respmsg = result + "，會是你最好的選擇。"
	default:
		respmsg = "你的意思是說，" + text + "。不過沒關係，一切都好起來的。"
	}

	photoBytes := imageByte("http://i1.hdslb.com/bfs/archive/ec20e0f5fbaa5f8252bf761e4b359edbb767fc43.png")
	// bot.SendText(chatID, respmsg)
	if _, ok := env.Nuniskill[chatID]; ok {
		bot.SendImageIdAddMsg(chatID, env.Nuniskill[chatID], respmsg)
	} else {
		bot.SendImageAddMsg(chatID, photoBytes, respmsg)
	}

}

//換
func (bot *Bot) killlakillnun(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	photo := *update.Message.Photo
	fId := photo[0].FileID
	env.Nuniskill[chatID] = fId
	bot.SendImageIdAddMsg(chatID, env.Nuniskill[chatID], "新修女上線，直到主教將原修女找回來")
}

//處理送往訊息管理群的訊息
func (bot *Bot) LogText(msg string) {
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
func (bot *Bot) RespNowTime() {
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
	default:

	}

}
