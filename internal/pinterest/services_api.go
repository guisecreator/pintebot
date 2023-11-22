package pinterest

import (
	"fmt"
	"github.com/carrot/go-pinterest/controllers"
	"github.com/carrot/go-pinterest/models"
)

// Pin API

func (p *PinterestServiceApi) GetPinsBySearch(
	tagName string,
) (*[]models.Pin, error) {
	pins, _, err := p.Client.Me.Search.Pins.Fetch(
		tagName,
		&controllers.MeSearchPinsFetchOptionals{
			Cursor: "some-cursor",
			Limit:  100,
		},
	)
	if err != nil {
		return nil, err
	}

	if len(*pins) == 0 {
		return nil, fmt.Errorf("no pins found")
	}

	return pins, nil
}

func (p *PinterestServiceApi) GetPinById(
	pinId string,
) (*models.Pin, error) {
	pin, err := p.Client.Pins.Fetch(pinId)
	if err != nil {
		return nil, err
	}

	noteOfPin := pin.Note
	if noteOfPin == "" {
		return nil, fmt.Errorf("pin fetch error: note: %s", pin)
	}

	return pin, nil
}

func (p *PinterestServiceApi) GetPinsFromBoard(
	boardSpec string,
) (*[]models.Pin, error) {
	pin, err := p.Client.Boards.Pins.Fetch(
		boardSpec,
		&controllers.BoardsPinsFetchOptionals{
			Cursor: "some-cursor",
		})
	if err != nil {
		return nil, err
	}

	return pin, nil
}

func (p *PinterestServiceApi) UpdatePin(
	pinSpec string,
) (*models.Pin, error) {
	update, err := p.Client.Pins.Update(
		pinSpec,
		&controllers.
			PinUpdateOptionals{
			Note: "test",
		})
	if err != nil {
		return nil, err
	}

	return update, nil
}

//Board API

func (p *PinterestServiceApi) GetBoard(
	boardId string,
) (*models.Board, error) {
	getBoard, err := p.Client.Boards.Fetch(boardId)
	if pinterestError, ok := err.(*models.PinterestError); ok {
		if pinterestError.StatusCode == 404 {
			_ = fmt.Errorf("board not found with id: %s", boardId)
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}

	noteOfBoard := getBoard.Counts.Pins
	if noteOfBoard == 0 {
		return nil, fmt.Errorf("board fetch error: note: %s", getBoard)
	}

	return getBoard, nil
}

func (p *PinterestServiceApi) UpdateBoard(
	boardSpec string,
) (*models.Board, error) {
	update, err := p.Client.Boards.Update(
		boardSpec,
		&controllers.
			BoardUpdateOptionals{
			Name:        "test",
			Description: "test",
		})
	if err != nil {
		return nil, err
	}

	return update, nil
}

//func (p *PinterestServiceApi) GetBoardsFromUserID(userId string) (*models.Board, error) {
//	user, err := p.Client.Users.Fetch(userId)
//	if err != nil {
//		return nil, err
//	}
//
//	boards := user.Counts.Boards
//	if boards == 0 {
//		return nil, fmt.Errorf("user fetch error: no boards found: %s", user)
//	}
//
//}
