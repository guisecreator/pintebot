package handlers

import (
	"context"
	"fmt"
	"github.com/guisecreator/pintebot/internal/config"
	"github.com/guisecreator/pintebot/internal/telegram/types"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"log"
)

type StartCommand struct {
	*types.CommandsOptions
}

func MessageError(userId telego.ChatID, message string, isReply bool) *telego.SendMessageParams {
	inlineKeyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(
				"Cancel",
				//сделать возврат на главное меню
			).WithCallbackData("start"),
		),
	)

	msgText := "Sorry, there was an error. Please try later."
	msg := tu.Message(
		userId,
		msgText,
	).WithReplyMarkup(inlineKeyboard).
		WithParseMode(telego.ModeHTML)

	return msg
}

func (start *StartCommand) HandleTelegramChatID(
	update telego.Update,
	ctx context.Context,
) (telego.ChatID, error) {
	chatID := tu.ID(
		update.
			Message.
			From.ID,
	)
	if update.Message != nil {
		return chatID, fmt.Errorf("message: %v", update.Message)
	}

	Callback := tu.ID(
		update.
			CallbackQuery.
			From.ID,
	)
	if update.CallbackQuery != nil {
		return Callback, fmt.Errorf("callback: %v", update.CallbackQuery)
	}

	return chatID, nil
}

func (start *StartCommand) HandleCancelStartCallback() th.Handler {
	return func(bot *telego.Bot, update telego.Update) {
	}
}

func (start *StartCommand) NewStartCommand() th.Handler {
	return func(bot *telego.Bot, update telego.Update) {
		//startCmd := &StartCommand{}
		//ctx := context.WithValue(context.Background(), "update", update)
		//chatID := startCmd.HandleTelegramChatID(update, ctx)
		//messageTextHello = messageTextHello + chatID.Error()

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
					WithCallbackData("find_pin_via_tag"),
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
					WithCallbackData("settings"),
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

		messageFilter := func(update *telego.Update) bool {
			return update.Message != nil
		}

		mainInlineKeyBoard := inlineKeyBoard
		if mainInlineKeyBoard != nil {
			messageFilter(&update)
		}

		if messageFilter(&update) {
			userName := update.Message.From.Username
			id := telego.ChatID{
				ID:       update.Message.Chat.ID,
				Username: userName,
			}

			msgTextHello := messages.Description
			if msgTextHello == "" {
				msgTextHello = "Hello! I'm PinteBot and my greeting is broken."
			}

			message := tu.Message(
				id,
				msgTextHello,
			).WithReplyMarkup(inlineKeyBoard).
				WithParseMode("HTML")

			_, sendMsgErr := bot.SendMessage(message)
			if err != nil {
				log.Printf("send message error: %v\n", sendMsgErr)
			}
		}
	}
}

func (start *StartCommand) HandleChangeIconNotification() {
	panic("implement me")
}
