package types

import (
	"github.com/guisecreator/pintebot/internal/config"
	"github.com/guisecreator/pintebot/internal/types"
)

type CommandsOptions struct {
	Services *types.BotServices
	Sessions SessionManager[Session]
	Commands *config.CommandsText
}

type MainMenu struct {
	CurrentPage int
	TotalPages  int
}
