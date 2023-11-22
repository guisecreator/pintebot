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
					KeyboardRow5.BackToStartButton,
			).
				WithCallbackData("cancel"),
		),
	)
	return inlineKeyboard
}

func (board *BoardsCommand) NewBoardCommand() th.Handler {
	return func(bot *telego.Bot, update telego.Update) {
		userId := tu.ID(update.CallbackQuery.From.ID)

		//inlineKeyboard := board.BuildKeyboard()
		//
		//messageText := "Board Commands:"
		//message := tu.Message(userId, messageText).
		//	WithReplyMarkup(inlineKeyboard).
		//	WithParseMode(telego.ModeHTML)

		_, botErr := bot.SendMessage(MessageError(userId, 23, "Command Error. Please try again later.", true))
		if botErr != nil {
			log.Printf("send message error: %v\n", botErr)
		}
	}
}
