package handlers

import "github.com/carrot/go-pinterest/models"

type BoardsByTitle struct {
}

func (board *BoardsByTitle) FindBoardByTitle() (*models.Board, error) {
	panic("implement me")
}
