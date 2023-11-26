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

type CommandsHandler struct {
	StartCommand  *StartCommand
	BoardsCommand *BoardsCommand
	TagsCommand   *TagsCommand
	HelpCommand   *HelpCommand
}

type StartCommand struct {
	*types.CommandsOptions
	logger *logrus.Logger
}

func MessageError(
	userId telego.ChatID,
	replyToMessageID int,
	message string,
	isReply bool,
) *telego.SendMessageParams {
	inlineKeyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(
				"Cancel",
			).WithCallbackData("cancel"),
		),
	)

	msg := tu.Message(
		userId,
		message,
	).WithReplyMarkup(inlineKeyboard).
		WithParseMode(telego.ModeHTML)

	if isReply {
		msg = msg.WithReplyToMessageID(replyToMessageID)
	}

	return msg
}

func BuildKeyboard() *telego.InlineKeyboardMarkup {
	messages, err := config.InitCommandsText("locales/en.yaml")
	if err != nil {
		log.Fatal(err)
	}

	inlineKeyBoard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(
				messages.StartCommand.InlineKeyboard.
					KeyboardRow1.FindPinViaTagButton,
			).
				WithCallbackData("find"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(
				messages.StartCommand.InlineKeyboard.
					KeyboardRow2.BoardsButton,
			).
				WithCallbackData("boards"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(
				messages.StartCommand.InlineKeyboard.
					KeyboardRow3.SettingsButton,
			).
				WithCallbackData("languages"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(
				messages.StartCommand.InlineKeyboard.
					KeyboardRow4.HelpButton,
			).
				WithCallbackData("help_info"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(
				messages.StartCommand.InlineKeyboard.
					KeyboardRow5.ProjectButton,
			).
				WithCallbackData("project_on_github").
				WithURL("https://github.com/guisecreator/pintebot"),
		),
	)
	return inlineKeyBoard
}

func (start *StartCommand) HandleStartCallback() th.Handler {
	return func(bot *telego.Bot, update telego.Update) {
		inlineKeyBoard := BuildKeyboard()

		callbackId := update.CallbackQuery.ID
		userId := tu.ID(update.CallbackQuery.From.ID)

		messages, err := config.InitCommandsText("locales/en.yaml")
		if err != nil {
			log.Fatal(err)
		}

		_, Boterr := bot.SendMessage(tu.Message(userId, messages.Description).
			WithReplyMarkup(inlineKeyBoard).
			WithParseMode(telego.ModeHTML))
		if err != nil {
			start.logger.Errorf("send message error: %v\n", Boterr)
		}

		callback := tu.CallbackQuery(callbackId)
		err = bot.AnswerCallbackQuery(callback)
		if err != nil {
			start.logger.Errorf("answer callback error: %v\n", err)
		}
	}
}

func (start *StartCommand) NewStartCommand() th.Handler {
	return func(bot *telego.Bot, update telego.Update) {
		inlineKeyBoard := BuildKeyboard()

		user_id := tu.ID(update.Message.From.ID)

		messages, err := config.InitCommandsText("locales/en.yaml")
		if err != nil {
			log.Fatal(err)
		}

		msgTextHello := messages.Description
		if msgTextHello == "" {
			msgTextHello = "Hello! I'm PinteBot and my greeting is broken."
		}

		message := tu.Message(
			user_id,
			msgTextHello,
		).WithReplyMarkup(inlineKeyBoard).
			WithParseMode(telego.ModeHTML)

		_, sendMsgErr := bot.SendMessage(message)
		if sendMsgErr != nil {
			log.Printf("send message error: %v\n", sendMsgErr)
		}
	}
}
