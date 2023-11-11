package telegram

import (
	"github.com/guisecreator/pintebot/internal/telegram/handlers"
	th "github.com/mymmrac/telego/telegohandler"
	"log"
)

func (service *TgBotService) handlersInit() error {
	handlers_init, err := handlers.NewCommandsHandler()
	if err != nil {
		log.Println(err)
	}
	// Start Bot.
	service.Handlers.Handle(
		handlers_init.
			StartCommand.
			NewStartCommand(),
		th.CommandEqual("start"),
	)
	// Redirect to main menu.
	service.Handlers.Handle(
		handlers_init.
			StartCommand.
			HandleStartCallback(),
		th.CallbackDataEqual("cancel"),
	)
	service.Handlers.Handle(
		handlers_init.
			BoardsCommand.
			NewBoardCommand(),
		th.CallbackDataEqual("boards"),
	)
	service.Handlers.Handle(
		handlers_init.
			TagsCommand.
			NewTagsCommand(),
		th.CallbackDataEqual("find_pin_via_tag"),
	)

	return nil
}
