package handlers

import (
	"github.com/guisecreator/pintebot/internal/config"
	"github.com/guisecreator/pintebot/internal/pinterest_service"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"log"
)

type PredicateService struct {
	ps *pinterest_service.PinterestService
}

func NewPredicateService(ps *pinterest_service.PinterestService) *PredicateService {
	return &PredicateService{
		ps: ps,
	}
}

func (p *PredicateService) PinterestServicePredicate(step string) th.Predicate {
	return func(update telego.Update) bool {
		if update.Message == nil || update.Message.From == nil || update.Message.Text == "" {
			log.Printf("UPDATE MSG NIL or TEXT is empty")
			return false
		}

		userId := update.Message.From.ID

		if userId == 0 || p.ps == nil {
			return false
		}

		if p.ps.PinterestMap == nil {
			p.ps.PinterestMap = make(map[int64]*pinterest_service.PinterestElement)
		}

		stepNow, ok := p.ps.PinterestMap[userId]
		if !ok || stepNow == nil {
			log.Printf("stepNow: %v", stepNow)
			return false
		}

		return stepNow.Step == step
	}
}

func (p *PredicateService) FirstMessage(step string) th.Predicate {
	return func(update telego.Update) bool {
		messages, err := config.InitCommandsText("locales/en.yaml")
		if err != nil {
			log.Fatal(err)
		}

		//pinsName, isPinsName := GetUserMessage(update)
		//if !isPinsName {
		//	tags.logger.Printf("get pins name error: %v\n", pinsName)
		//	return
		//}
		userId := tu.ID(update.Message.From.ID)

		messageText := messages.AnyTagText
		message := tu.Message(
			userId,
			messageText,
		).WithParseMode(telego.ModeHTML)

		bot := telego.Bot{}

		_, botErr := bot.SendMessage(message)
		if botErr != nil {
			log.Printf("send message error: %v\n", botErr)
		}
		return true
	}
}
