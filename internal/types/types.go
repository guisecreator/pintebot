package types

import (
	"github.com/guisecreator/pintebot/internal/config"
	"github.com/guisecreator/pintebot/internal/pinterest"
	"github.com/guisecreator/pintebot/internal/pinterest/api"
)

type BotServices struct {
	Config    *config.Config
	Pinterest *pinterest.PinterestServiceApi
	Boards    *api.Board
	Tags      *api.Pin
}
