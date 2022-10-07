package controller

import (
	"fmt"
	"testproject/confession/env"
	"testproject/confession/modules/tgbot"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/robfig/cron/v3"
)

func OnBot() {
	var telegbot *tgbot.Bot
	telegbot = tgbot.New(env.TGToken)

	telegbot.LogText("您的修女大人已上線摟")
	// 報時排程
	go timekeepSchedule(telegbot)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updates, _ := telegbot.Botapi.GetUpdatesChan(updateConfig)

	// go runddd(bot)
	for update := range updates {
		go telegbot.HandleUpdate(update)
	}
}

func timekeepSchedule(telegbot *tgbot.Bot) {
	c := cron.New()
	id, _ := c.AddFunc("0 9 * * 1-5", func() {
		telegbot.RespNowTime()
	})
	fmt.Println(id)
	id, _ = c.AddFunc("0 12 * * 1-5", func() {
		telegbot.RespNowTime()
	})
	fmt.Println(id)
	id, _ = c.AddFunc("0 18 * * 1-5", func() {
		telegbot.RespNowTime()
	})
	fmt.Println(id)
	id, _ = c.AddFunc("0 22 * * 1-5", func() {
		telegbot.RespNowTime()
	})
	fmt.Println(id)
	// id, _ = c.AddFunc("5,6,7,8,9 * * * 1-5", func() {
	// 	telegbot.TestRespNowTime()
	// })
	// fmt.Println(id)
	c.Start()
}
