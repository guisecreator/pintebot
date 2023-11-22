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

	//ps := service.BotServices.PinterestService
	//predicate := handlers.NewPredicateService(ps)

	// Start Bot.
	service.Handlers.Handle(
		handlers_init.StartCommand.NewStartCommand(),
		th.CommandEqual("start"),
	)
	// Redirect to main menu.
	service.Handlers.Handle(
		handlers_init.StartCommand.HandleStartCallback(),
		th.CallbackDataEqual("cancel"),
	)

	service.Handlers.Handle(
		handlers_init.BoardsCommand.NewBoardCommand(),
		th.CallbackDataEqual("boards"),
	)

	service.Handlers.Handle(
		handlers_init.TagsCommand.MessageTag(),
		th.AnyCallbackQueryWithMessage(),
		th.CallbackDataEqual("find_pin_via_tag"),
	)

	service.Handlers.Handle(
		handlers_init.TagsCommand.NewTagsCommand(),
		th.AnyMessage(),
	)

	//Help information command
	service.Handlers.Handle(
		handlers_init.HelpCommand.HelpCommand(),
		th.CallbackDataEqual("help_info"),
	)

	return nil
}
