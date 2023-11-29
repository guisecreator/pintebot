package types

import (
	"github.com/guisecreator/pintebot/internal/config"
	"github.com/guisecreator/pintebot/internal/db/models"
	"github.com/guisecreator/pintebot/internal/message_handler"
	"github.com/guisecreator/pintebot/internal/pics_service"
	"github.com/guisecreator/pintebot/internal/pinterest"
)

type BotServices struct {
	Config          *config.Config
	Tags            *models.Tag
	MessageHandler  message_handler.MessageHandler
	PinterestClient *pinterest.Client
	PicsService     *pics_service.PicsService
}
