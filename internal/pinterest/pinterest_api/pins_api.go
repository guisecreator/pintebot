package pinterest_api

import (
	"fmt"
	"github.com/carrot/go-pinterest"
	"github.com/carrot/go-pinterest/controllers"
	"github.com/carrot/go-pinterest/models"
)

type Pin struct {
	Controller *controllers.PinsController
	PinModel   *models.Pin
	PinCounts  *models.PinCounts
}

func GetPinsBySearch(
	client pinterest.Client,
	tagName string,
) (*[]models.Pin, error) {
	pins, page, err := client.Me.Search.Pins.Fetch(
		tagName,
		&controllers.MeSearchPinsFetchOptionals{
			Cursor: "some-cursor",
			Limit:  15,
		},
	)
	if pinterestError, ok := err.(*models.PinterestError); ok {
		if pinterestError.StatusCode == 404 {
			_ = fmt.Errorf("pin fetch error: %s", pins)
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}

	user, userErr := client.Users.Fetch((*pins)[0].Creator.Id)
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

func (p *Pin) GetPinById(client pinterest.Client, pinId string) (*models.Pin, error) {
	pin, err := client.Pins.Fetch(pinId)
	if err != nil {
		return nil, err
	}

	return pin, nil
}

func (p *Pin) GetPinsByIds(client pinterest.Client, pinIds []string) (*[]models.Pin, error) {
	//pins := make([]models.Pin, 0)

	return nil, nil
}
