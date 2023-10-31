package handlers

import (
	"github.com/guisecreator/pintebot/internal/config"
	"github.com/guisecreator/pintebot/internal/telegram/types"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/sirupsen/logrus"
	"log"
)

type TagsCommand struct {
	*types.CommandsOptions
	logger *logrus.Logger
}

func (tags *TagsCommand) NewTagsCommand() th.Handler {
	return func(bot *telego.Bot, update telego.Update) {
		messages, err := config.InitCommandsText("locales/en.yaml")
		if err != nil {
			log.Fatal(err)
		}

		userId := tu.ID(
			update.
				CallbackQuery.
				From.ID)
		if update.Message != nil && update.Message.Text != "" {
			//TODO: userText handler
			//userText := update.Message.Text

			inlineKeyboard := tu.Keyboard(
				tu.KeyboardRow(
					tu.KeyboardButton(
						messages.TagsCommand.InlineKeyboard.
							KeyboardRow1.NextButton,
					),
				),
			).WithResizeKeyboard()

			buttonMessage := tu.Message(
				userId,
				"",
			).WithReplyMarkup(inlineKeyboard)
			_, buttonErr := bot.SendMessage(buttonMessage)
			if buttonErr != nil {
				log.Printf("send button error: %v\n", buttonErr)
			}

		} else {
			messageText := messages.AnyTagText
			message := tu.Message(
				userId,
				messageText,
			).WithParseMode(telego.ModeHTML)

			_, botErr := bot.SendMessage(message.WithProtectContent())
			if botErr != nil {
				log.Printf("send message error: %v\n", botErr)
			}
		}
	}
}
