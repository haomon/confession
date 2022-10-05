package env

import (
	"strconv"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Load() {
	setLog()
	SetDefultValue()
}

func setLog() {
	// Use Zap Production Config but change time format to ISO8601
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	// Default setting (to Production)
	customConfig := &zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       false,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
		DisableStacktrace: true,
	}
	customConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	customConfig.Development = true
	l, err := customConfig.Build()
	if err != nil {
		panic(err)
	}
	defer l.Sync()

	// Config zap logger
	//cfg := zap.NewProductionConfig()
	//cfg.DisableStacktrace = true
	//cfg.Sampling = nil
	//cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	//l, _ := cfg.Build()
	zap.ReplaceGlobals(l)
}

func SetDefultValue() {
	AlertTitle = "CaAlert通知"
	ChatPrefix = "ㄡ~修女啊，"
}
func GetchatIDs(str_chatIDs []string) (err error, chatIDs []int64) {
	for _, str_chatID := range str_chatIDs {
		chatID, err := strconv.ParseInt(str_chatID, 10, 64)
		if err != nil {
			break
		}
		chatIDs = append(chatIDs, chatID)
	}
	return err, chatIDs
}

// func getMustString(key string) string {
// 	bot, err = tgbotapi.NewBotAPI(youToken)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	bot.Debug = false

// 	return v
// }

// func getMustInt(key string) int {
// 	v, err := config.GetInt(key)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return v
// }

// func getMustStringSlice(key string) []string {
// 	v, err := config.GetStringSlice(key)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return v
// }

// func getMustSectionString(section string, key string) string {
// 	v, err := config.GetSectionConfig(section)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return v[key].(string)
// }
