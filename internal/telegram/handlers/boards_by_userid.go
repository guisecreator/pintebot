package handlers

import (
	"github.com/carrot/go-pinterest/models"
	"github.com/guisecreator/pintebot/internal/telegram/types"
)

type BoardsByAccount struct {
	*types.CommandsOptions
}

func (board *BoardsCommand) FindBoardByUserName() (*models.Board, error) {
	panic("implement me")
}
