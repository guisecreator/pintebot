package pinterest

import (
	"fmt"
	"github.com/carrot/go-pinterest/controllers"
	"github.com/carrot/go-pinterest/models"
)

// TODO: max кол-во запросов по пинам 30. то есть бот получает только 30 пинов за 1 запрос по тегу
func (p *PinterestServiceApi) GetPinsBySearch(tagName string) (*[]models.Pin, error) {
	pins, page, err := p.Client.Me.Search.Pins.Fetch(
		tagName,
		&controllers.MeSearchPinsFetchOptionals{
			Cursor: "some-cursor",
			Limit:  30,
		},
	)
	if err != nil {
		return nil, err
	}

	user, userErr := p.Client.Users.Fetch((*pins)[0].Creator.Id)
	if userErr != nil {
		return nil, userErr
	}

	for _, i := range page.Next {
		if i == int32(page.Next[len(page.Next)-1]) {
			return nil, fmt.Errorf("pin fetch error: %s", pins)
		}
		pin := models.Pin{
			Id: string(i),
			Creator: models.Creator{
				Url:       user.Url,
				Id:        user.Id,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			},
		}
		*pins = append(*pins, pin)

		return pins, nil
	}

	return pins, nil
}

func (p *PinterestServiceApi) GetPinById(pinId string) (*models.Pin, error) {
	pin, err := p.Client.Pins.Fetch(pinId)
	if err != nil {
		return nil, err
	}

	return pin, nil
}

func (p *PinterestServiceApi) GetPinsByIds(pinIds []string) (*[]models.Pin, error) {
	panic("implement me")
}
