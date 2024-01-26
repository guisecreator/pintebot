package telegram

import (
	"github.com/guisecreator/pintebot/internal/telegram/handlers"
	th "github.com/mymmrac/telego/telegohandler"
)

func (service *TgBotService) handlersInit() error {
	handlersInit := handlers.CommandsHandler{}

	ps := service.BotServices.PicsService
	predicate := handlers.NewPredicateService(ps)

	service.Handlers.Handle(
		handlersInit.StartCommand.NewStartCommand,
		th.CommandEqual("start"),
	)

	//Login command handler
	service.Handlers.Handle(
		handlersInit.StartCommand.SendLoginUrl,
		th.TextEqual("/login"),
	)

	service.Handlers.Handle(
		handlersInit.StartCommand.HandleStartCallback,
		th.CallbackDataEqual("cancel"),
	)

	service.Handlers.Handle(
		handlersInit.BoardsCommand.NewBoardCommand,
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

	return nil
}
