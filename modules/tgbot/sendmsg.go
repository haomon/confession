package tgbot

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
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

	userchatID := update.Message.Chat.ID
	// fmt.Println(userchatID)
	// fmt.Println(chatID)
	switch {
	case text != "":
		if env.NXNmode {
			if text == "關閉" {
				env.NXNmode = false
				return
			}
			if _, ok := env.NXN[text]; ok {
				bot.SendText(userchatID, env.NXN[text])
				return
			}
		}
		if strings.HasPrefix(text, env.ChatPrefix) {
			bot.AnalyzeText(update)
		} else {
			bot.MemeText(update)
		}
	case sticker != nil:
		fmt.Println(sticker.FileID)
		// replyMsg := tgbotapi.NewStickerShare(env.LogchatID, "CAACAgUAAxkBAAOUYzpShhs4JzzEK5cEDqSe0e1KzeYAAmwBAAJvL7ojW2D8vo04eqwqBA")
		// replyMsg.ReplyToMessageID = update.Message.MessageID
		// _, _ = bot.Send(replyMsg)
	case photo != nil:
		// fmt.Println(update.Message.Caption)
		caption := update.Message.Caption
		if len(caption) > 0 && caption == "安妮亞喜歡這個修女" {
			bot.killlakillnun(update)
		}
	}

}

func c8763(text string) (result int) {
	result = -1
	c8763 := []string{"Starburst", "Stream", "スターバースト", "ストリーム", "星爆氣流斬", "星光連流擊", "C8763", "還要更快", "撐十秒", "撐１０秒", "稱十秒", "稱１０秒", "克萊因拜託了"}
	for k, v := range c8763 {
		if strings.Index(text, v) >= 0 {
			result = k
			break
		}
	}
	return result
}

// 對特定含埂訊息做回覆
func (bot *Bot) MemeText(update tgbotapi.Update) {
	text := update.Message.Text
	chatID := update.Message.Chat.ID
	switch {
	case strings.Index(text, "請你喝飲料") >= 0:
		photoBytes := imageByte("https://pbs.twimg.com/media/C2YVjnMXEAMhzYW.jpg")
		respmsg := "機...油...好難...喝..."
		bot.SendImageAddMsg(chatID, photoBytes, respmsg)
		// bot.SendText(chatID, respmsg)
	case strings.Index(text, "我看看") >= 0 || strings.Index(text, "讓我康康") >= 0:
		replyMsg := tgbotapi.NewStickerShare(chatID, "CAACAgUAAxkBAAIB9WNFF3yOpa4r1_a51cClVpIBepenAAI6AAMUHVkORir3PjXLGrkqBA")
		_, _ = bot.Botapi.Send(replyMsg)
	case strings.Index(text, "知道你家住哪") >= 0:
		replyMsg := tgbotapi.NewStickerShare(chatID, "CAACAgUAAxkBAAICG2NFIlW80eyZ9jmv6HiOjUC9cItzAAJBAAMUHVkO8TQmvtnIRl8qBA")
		_, _ = bot.Botapi.Send(replyMsg)
	case strings.Index(text, "不懂喔") >= 0:
		replyMsg := tgbotapi.NewStickerShare(chatID, "CAACAgUAAxkBAAICH2NFI3K6LeRYy31I8Zt_m7SNytH3AAIxAAMUHVkOHqidqI8hB-sqBA")
		_, _ = bot.Botapi.Send(replyMsg)
	case strings.Index(text, "超勇") >= 0:
		replyMsg := tgbotapi.NewStickerShare(chatID, "CAACAgUAAxkBAAICFmNFId-wtXYqE4G20Rf2_ok-DhmYAAItAAMUHVkOjc0in88yQpYqBA")
		_, _ = bot.Botapi.Send(replyMsg)
	case strings.Index(text, "身材不錯") >= 0 && strings.Index(text, "結實") >= 0:
		replyMsg := tgbotapi.NewStickerShare(chatID, "CAACAgUAAxkBAAICF2NFIfCUERwudxDH2xVAcFrzYqD7AAIuAAMUHVkOqgTwd4XPjwMqBA")
		_, _ = bot.Botapi.Send(replyMsg)
	case strings.Index(text, "幹嘛啦") >= 0:
		replyMsg := tgbotapi.NewStickerShare(chatID, "CAACAgUAAxkBAAICQGNFQlT3ZQm9FvBO7SqPyDkWeEJXAAIyAAMUHVkOpcKV2-RN58cqBA")
		_, _ = bot.Botapi.Send(replyMsg)
	case strings.Index(text, "新遊戲") >= 0:
		replyMsg := tgbotapi.NewStickerShare(chatID, "CAACAgUAAxkBAAICHmNFIqjK2K1GVfxWRkvOv4OIK3CuAAI0AAMUHVkOe_g-VDu-rcAqBA")
		_, _ = bot.Botapi.Send(replyMsg)
		replyMsg = tgbotapi.NewStickerShare(chatID, "CAACAgUAAxkBAAICGWNFIiJ_gG31qwmOgvI1WdfnbZy4AAI1AAMUHVkOQTkl7mX3YFAqBA")
		_, _ = bot.Botapi.Send(replyMsg)
	case strings.Index(text, "很勇喔") >= 0:
		replyMsg := tgbotapi.NewStickerShare(chatID, "CAACAgUAAxkBAAICFWNFIavZtqrafX0V4Xd66zf0Bg_RAAIrAAMUHVkOUgM9EI4DRLMqBA")
		_, _ = bot.Botapi.Send(replyMsg)
	case strings.Index(text, "94遜") >= 0 || strings.Index(text, "就是遜") >= 0:
		replyMsg := tgbotapi.NewStickerShare(chatID, "CAACAgUAAxkBAAICMGNFP6DjWqi3Z3EaurP6LIBgFqVMAAIqAAMUHVkO6tFXmX0qa1gqBA")
		_, _ = bot.Botapi.Send(replyMsg)
	case strings.Index(text, "沒錢") >= 0 || strings.Index(text, "沒鉗") >= 0:
		replyMsg := tgbotapi.NewStickerShare(chatID, "CAACAgUAAxkBAAIB9mNFGDWYpnPDymK-c3VI-TICBC8TAALIAAMBVzELE8AOW25TLAUqBA")
		_, _ = bot.Botapi.Send(replyMsg)
	case strings.Index(text, "色") >= 0 || strings.Index(text, "瑟") >= 0:
		replyMsg := tgbotapi.NewStickerShare(chatID, "CAACAgUAAxkBAAIB-WNFGPF4FzW2JfYkWu3P8uSkLQgVAAIrAgACuZoSHXnHOjczz0LfKgQ")
		_, _ = bot.Botapi.Send(replyMsg)
	case strings.Index(text, "射了") >= 0:
		replyMsg := tgbotapi.NewStickerShare(chatID, "CAACAgUAAxkBAAIB-2NFGev0YXewnn_7IT2WgRy7cbHwAAIWAAO3cQkH5abETHN2kOgqBA")
		_, _ = bot.Botapi.Send(replyMsg)
	case c8763(text) >= 0:
		photoBytes := imageByte("https://truth.bahamut.com.tw/s01/201609/47eb58a327c095b5a80512e8e4720bf3.PNG")
		bot.SendImage(chatID, photoBytes)
		// 逮捕派大星 CAACAgUAAxkBAAIB92NFGHwjslwqHnsbSI7g63836EDTAALHAAMBVzELUJbnWU36zBoqBA
		// 準備好的派大星 CAACAgUAAxkBAAIB-GNFGKXtRvlfFjfgXg9qh--wEtXNAALFAAMBVzELDTLKpBMD8loqBA
		// CAACAgUAAxkBAAICFWNFIavZtqrafX0V4Xd66zf0Bg_RAAIrAAMUHVkOUgM9EI4DRLMqBA 超勇
		// CAACAgUAAxkBAAICFmNFId-wtXYqE4G20Rf2_ok-DhmYAAItAAMUHVkOjc0in88yQpYqBA 身材不錯
		// 還可以教你轉大人 CAACAgUAAxkBAAICGWNFIiJ_gG31qwmOgvI1WdfnbZy4AAI1AAMUHVkOQTkl7mX3YFAqBA
		// 連紅啦 CAACAgUAAxkBAAICGmNFIjpLWtypezYPE_-WocpqERUNAAI3AAMUHVkOmY8_25UFVp0qBA
		// 想你 CAACAgUAAxkBAAICG2NFIlW80eyZ9jmv6HiOjUC9cItzAAJBAAMUHVkO8TQmvtnIRl8qBA
		// 比遊戲還次際 CAACAgUAAxkBAAICHmNFIqjK2K1GVfxWRkvOv4OIK3CuAAI0AAMUHVkOe_g-VDu-rcAqBA
		// CAACAgUAAxkBAAICGGNFIhFSTy8kvxdWhyZJ4cKMnnohAAIvAAMUHVkOQ64sh_dDBlkqBA 都幾歲了
		// CAACAgUAAxkBAAICF2NFIfCUERwudxDH2xVAcFrzYqD7AAIuAAMUHVkOqgTwd4XPjwMqBA 你幹嘛啦
	// case strings.Index(text, "請你喝飲料") >= 0:

	// 	photoBytes := imageByte("https://pbs.twimg.com/media/C2YVjnMXEAMhzYW.jpg")
	// 	bot.SendImage(chatID, photoBytes)
	default:
	}

}

//分析 「ㄡ~修女啊，」 後該做什麼
func (bot *Bot) AnalyzeText(update tgbotapi.Update) {

	var lotton bool
	var respmsg string
	text := update.Message.Text
	chatID := update.Message.Chat.ID
	text = text[len(env.ChatPrefix):]

	if strings.Contains(text, "[") && strings.Contains(text, "]") {
		lotton = true
	}

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
	case text == "九九":
		env.NXNmode = true
		respmsg = "無情乘法表"
		bot.SendText(chatID, respmsg)
		return
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
	bot.LogText("報時執行小時:" + strconv.Itoa(time.Now().Hour()))
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
		respmsg := "機...油...好難...喝..."
		bot.SendText(env.ChatID, respmsg)
	}

}

//不用不用報時
func (bot *Bot) TestRespNowTime() {
	fridayTen := "CAACAgUAAxkBAAOfYzpUZR-HubpnmgABHGePT8ayrnvjAAJVBQACby-6I5UH8HAxbvwCKgQ"
	fmt.Println(time.Now().Hour())
	switch time.Now().Hour() {
	default:
		replyMsg := tgbotapi.NewStickerShare(env.LogchatID, fridayTen)
		_, err := bot.Botapi.Send(replyMsg)
		if err != nil {
			bot.LogText(err.Error())
		}
	}

}

// // 星報梗
// func c8763(text string) (result int) {
// 	c8763 := map[string]string{
// 		"Starburst": "",
// 		"Stream":    "", "スターバースト": "", "ストリーム": "", "星爆氣流斬": "", "星光連流擊": "", "C8763": "", "星爆": "", "西瓜榴槤": "", "還要更快": "", "撐十秒": "", "撐１０秒": "", "稱十秒": "", "稱１０秒": "", "克萊因拜託了": ""}
// 	for k, v := range c8763 {
// 		if strings.Index(text, v) >= 0 {
// 			result = k
// 			break
// 		}
// 	}
// 	return result
// }
