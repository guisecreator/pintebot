package pinterest

import (
	"github.com/carrot/go-pinterest"
	"github.com/carrot/go-pinterest/models"
)

type Interface interface {
	GetBoardsFromUserID(userId string) (*models.Board, error)
	GetBoard(boardId string) (*models.Board, error)

	UpdateBoard(boardSpec string) (*models.Board, error)

	GetPinsBySearch(client pinterest.Client, tagName string) (*[]models.Pin, error)
	GetPinById(client pinterest.Client, pinId string) (*models.Pin, error)
	GetPinsFromBoard(boardSpec string) (*[]models.Pin, error)
}
