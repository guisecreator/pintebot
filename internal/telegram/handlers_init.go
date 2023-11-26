package telegram

import (
	"github.com/guisecreator/pintebot/internal/telegram/handlers"
	th "github.com/mymmrac/telego/telegohandler"
)

func (service *TgBotService) handlersInit() error {
	handlersInit :=  handlers.CommandsHandler{}

	ps := service.BotServices.PinterestService
	predicate := handlers.NewPredicateService(ps)

	service.Handlers.Handle(
		handlersInit.StartCommand.NewStartCommand(),
		th.CommandEqual("start"),
	)

	service.Handlers.Handle(
		handlersInit.StartCommand.HandleStartCallback(),
		th.CallbackDataEqual("cancel"),
	)

	service.Handlers.Handle(
		handlersInit.BoardsCommand.NewBoardCommand(),
		th.CallbackDataEqual("boards"),
	)

	service.Handlers.Handle(
		handlersInit.TagsCommand.MessageTag,
		th.TextEqual("/find_pins"),
	)

	service.Handlers.Handle(
		handlersInit.TagsCommand.NewTagsCommand,
		predicate.NewTagsPredicate,
	)

	service.Handlers.Handle(
		handlersInit.HelpCommand.HelpCommand(),
		th.CallbackDataEqual("help_info"),
	)

	return nil
}
