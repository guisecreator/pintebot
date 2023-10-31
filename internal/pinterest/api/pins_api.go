package api

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

func (p *Pin) GetPinsBySearch(
	client pinterest.Client,
	tagName string,
) (*[]models.Pin, error) {
	pins, _, err := client.Me.Search.Pins.Fetch(
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

	//for _, i := range page.Next {
	//	pins = append(pins, page.Cursor[i])
	//	return pins, nil
	//}

	return pins, nil
}

func (p *Pin) GetPin(client pinterest.Client, pinId string) (*models.Pin, error) {
	pin, err := client.Pins.Fetch(pinId)
	if pinterestError, ok := err.(*models.PinterestError); ok {
		if pinterestError.StatusCode == 404 {
			_ = fmt.Errorf("pin fetch error: %s", pin)
		} else {
			return nil, err
		}
	}

	return pin, nil
}

func (p *Pin) GetPins(client pinterest.Client, pinIds []string) (*[]models.Pin, error) {
	//pins := make([]models.Pin, 0)

	return nil, nil
}
