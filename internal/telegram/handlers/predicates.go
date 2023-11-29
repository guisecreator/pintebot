package handlers

import (
	"github.com/guisecreator/pintebot/internal/config"
	"github.com/guisecreator/pintebot/internal/pics_service"
	"github.com/mymmrac/telego"
	"log"
)

type PredicateService struct {
	ps *pics_service.PicsService
}

func NewPredicateService(ps *pics_service.PicsService) *PredicateService {
	return &PredicateService{
		ps: ps,
	}
}

func (p *PredicateService) NewTagsPredicate(update telego.Update) bool {
	messages, err := config.InitCommandsText("locales/en.yaml")
	if err != nil {
		log.Fatal(err)
	}

	if update.Message != nil &&
		update.Message.ReplyToMessage != nil &&
		update.Message.ReplyToMessage.Text != "" {
		return update.Message.ReplyToMessage.Text == messages.AnyTagText
	}
	return false
}
