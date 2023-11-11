package handlers

import (
	"errors"
	"github.com/guisecreator/pintebot/internal/config"
	"github.com/guisecreator/pintebot/internal/telegram/types"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"log"
)

type BoardsCommand struct {
	*types.CommandsOptions
}

var (
	pinterestInvalidData = "invalid login names in request"
	invalidNameError     = errors.New(pinterestInvalidData)
)

func (board *BoardsCommand) BuildKeyboard() *telego.InlineKeyboardMarkup {
	messages, err := config.InitCommandsText("locales/en.yaml")
	if err != nil {
		log.Fatal(err)
	}

	inlineKeyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(
				messages.BoardsCommand.InlineKeyboard.
					KeyboardRow1.SpecificUsersBoardsButton,
			).
				WithCallbackData("fetch_user_boards"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(
				messages.BoardsCommand.InlineKeyboard.
					KeyboardRow2.BoardsByTitleButton,
			).
				WithCallbackData("pinterest_boards_by_title"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(
				messages.BoardsCommand.InlineKeyboard.
					KeyboardRow3.HistoryOfBoardsButton,
			).
				WithCallbackData("history_boards"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(
				messages.BoardsCommand.InlineKeyboard.
					KeyboardRow4.HelpButton,
			).
				WithCallbackData("back_to_main_menu"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(
				messages.BoardsCommand.InlineKeyboard.
					KeyboardRow5.BackToStartButton,
			).
				WithCallbackData("cancel"),
		),
	)
	return inlineKeyboard
}

func (board *BoardsCommand) NewBoardCommand() th.Handler {
	return func(bot *telego.Bot, update telego.Update) {
		inlineKeyboard := board.BuildKeyboard()

		userId := tu.ID(update.CallbackQuery.From.ID)
		messageText := "Board Commands:"
		message := tu.Message(userId, messageText).
			WithReplyMarkup(inlineKeyboard).
			WithParseMode(telego.ModeHTML)

		_, botErr := bot.SendMessage(message)
		if botErr != nil {
			log.Printf("send message error: %v\n", botErr)
		}
	}
}
