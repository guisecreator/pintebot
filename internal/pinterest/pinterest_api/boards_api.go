package pinterest_api

import (
	"fmt"
	"github.com/BrandonRomano/wrecker"
	"github.com/carrot/go-pinterest/controllers"
	"github.com/carrot/go-pinterest/models"
	"net/http"
)

type Board struct {
	Controller  *controllers.BoardsController
	BoardModel  *models.Board
	BoardCounts *models.BoardCounts
}

func NewBoards() (Board, error) {
	contentType := "application/json"

	wc := &wrecker.Wrecker{
		BaseURL:    "https://api.pinterest.com/v1/",
		HttpClient: &http.Client{
			//Transport: &http.Transport{},
			//CheckRedirect: func(req *http.Request,
			//	via []*http.Request) error {
			//	return http.ErrUseLastResponse
			//},
		},
		DefaultContentType: contentType,
		RequestInterceptor: func(request *wrecker.Request) error {
			return nil
		},
	}

	return Board{
		Controller: controllers.NewBoardsController(wc),
	}, nil
}

func (b *Board) GetBoardsFromUserID(userId string) (*models.Board, error) {
	panic("implement me")
}

func (b *Board) GetBoards(boardIds, board []string) ([]*models.Board, error) {
	panic("implement me")
}

func (b *Board) GetBoard(boardId string) (*models.Board, error) {
	getBoard, err := b.Controller.Fetch(boardId)
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

func (b *Board) UpdateBoard(boardSpec string) (*models.Board, error) {
	updateOptionals := controllers.
		BoardUpdateOptionals{
		Name:        "test",
		Description: "test",
	}

	update, err := b.Controller.Update(boardSpec, &updateOptionals)
	if err != nil {
		return nil, err
	}

	return update, nil
}

func (b *Board) DeleteBoard(boardId string) (bool, error) {
	err := b.Controller.Delete(boardId)
	if err != nil {
		return false, fmt.Errorf("error deleting board: %w", err)
	}

	return true, nil
}
