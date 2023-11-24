package telegram

import (
	"github.com/guisecreator/pintebot/internal/telegram/handlers"
	th "github.com/mymmrac/telego/telegohandler"
)

func (service *TgBotService) handlersInit() error {
	handlersInit :=  handlers.CommandsHandler{}

	//ps := service.BotServices.PinterestService
	//predicate := handlers.NewPredicateService(ps)

	// Start Bot.
	service.Handlers.Handle(
		handlersInit.StartCommand.NewStartCommand(),
		th.CommandEqual("start"),
	)
	// Redirect to main menu.
	service.Handlers.Handle(
		handlersInit.StartCommand.HandleStartCallback(),
		th.CallbackDataEqual("cancel"),
	)

	//Unknow command
	service.Handlers.Handle(
		handlersInit.StartCommand.NotSupportedCommand(),
		th.AnyMessage(),
	)

	service.Handlers.Handle(
		handlersInit.BoardsCommand.NewBoardCommand(),
		th.CallbackDataEqual("boards"),
	)

	service.Handlers.Handle(
		handlersInit.TagsCommand.MessageTag(),
		th.AnyCallbackQueryWithMessage(),
		th.CallbackDataEqual("find_pin_via_tag"),
	)

	service.Handlers.Handle(
		handlersInit.TagsCommand.NewTagsCommand(),
		th.TextEqual("find_pins"),
	)

	//Help information command
	service.Handlers.Handle(
		handlersInit.HelpCommand.HelpCommand(),
		th.CallbackDataEqual("help_info"),
	)

	return nil
}
