package types

import (
	"github.com/guisecreator/pintebot/internal/config"
	"github.com/guisecreator/pintebot/internal/db/models"
	"github.com/guisecreator/pintebot/internal/pinterest"
	"github.com/guisecreator/pintebot/internal/pinterest_service"
)

type BotServices struct {
	Config       *config.Config
	PinterestAPI *pinterest.PinterestServiceApi
	//Boards       *pinterest_api.Board
	//Tags         *pinterest_api.Pin
	Tags  *models.Tag
	Users *models.User

	PinterestService *pinterest_service.PinterestService
}
