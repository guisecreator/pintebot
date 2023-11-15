package pinterest

import (
	"fmt"
	"github.com/carrot/go-pinterest/controllers"
	"github.com/carrot/go-pinterest/models"
)

func (p *PinterestServiceApi) GetBoardsFromUserID(userId string) (*models.Board, error) {
	panic("implement me")
}

func (p *PinterestServiceApi) GetBoards(boardIds, board []string) ([]*models.Board, error) {
	panic("implement me")
}

func (p *PinterestServiceApi) GetBoard(boardId string) (*models.Board, error) {
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

	return getBoard, nil
}

func (p *PinterestServiceApi) UpdateBoard(boardSpec string) (*models.Board, error) {
	updateOptionals := controllers.
	BoardUpdateOptionals{
		Name:        "test",
		Description: "test",
	}

	update, err := p.Client.Boards.Update(boardSpec, &updateOptionals)
	if err != nil {
		return nil, err
	}

	return update, nil
}
