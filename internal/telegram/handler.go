package telegram

import (
	"github.com/guisecreator/pintebot/internal/telegram/handlers"
	th "github.com/mymmrac/telego/telegohandler"
	"log"
)

func (service *TgBotService) initHandlers() error {
	initHandler, err := handlers.NewCommandsHandler()
	if err != nil {
		log.Println(err)
	}

	service.Handlers.Handle(
		initHandler.
			StartCommand.
			NewStartCommand(),
		th.CommandEqual("start"),
	)
	service.Handlers.Handle(
		initHandler.
			BoardsCommand.
			NewBoardCommand(),
		th.CallbackDataEqual("boards"),
	)
	service.Handlers.Handle(
		initHandler.
			TagsCommand.
			NewTagsCommand(),
		th.CallbackDataEqual("find_pin_via_tag"),
	)

	return nil
}
