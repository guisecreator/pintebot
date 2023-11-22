package handlers

import (
	"github.com/guisecreator/pintebot/internal/config"
	"github.com/guisecreator/pintebot/internal/telegram/types"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"log"
)

type HelpCommand struct {
	*types.CommandsOptions
}

func (help *HelpCommand) HelpCommand() th.Handler {
	return func(bot *telego.Bot, update telego.Update) {
		messages, err := config.InitCommandsText("locales/en.yaml")
		if err != nil {
			log.Fatal(err)
		}

		userId := tu.ID(update.CallbackQuery.From.ID)

		_, botErr := bot.SendMessage(MessageError(userId, 1, messages.HelpInfoText, true))
		if botErr != nil {
			log.Printf("send message error: %v\n", botErr)
		}
	}
}
