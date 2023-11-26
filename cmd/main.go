package main

import (
	"github.com/guisecreator/pintebot/internal/config"
	"github.com/guisecreator/pintebot/internal/telegram"
	"github.com/guisecreator/pintebot/internal/types"
	"github.com/guisecreator/pintebot/pkg/logger"
)

func main() {
	logg := logger.InitLogger()

	cfg, err := config.NewConfig()
	if err != nil {
		logg.Panicf("config init: %v", err)
	}

	_, err = config.InitCommandsText("locales/en.yaml")
	if err != nil {
		logg.Panicf("commands init: %v", err)
	}

	// pinterestService, err := pinterest.NewPinterestService(*cfg)
	// if err != nil {
	// 	logg.Panicf("pinterest service init: %v", err)
	// }

	services := &types.BotServices{
		Config: cfg,
		// PinterestAPI: pinterestService,
	}

	bot, err := telegram.NewTelegram(*cfg, *services, cfg.TgToken, logg)
	if err != nil {
		logg.Panicf("telegram service init: %v", err)
	}

	bot.StartService()
}
