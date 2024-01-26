package handlers

import (
	"github.com/guisecreator/pintebot/internal/telegram/types"
	"github.com/mymmrac/telego"
)

type Auth struct {
	Token string
}

type AuthCommand struct {
	*types.CommandsOptions
}

func (a *AuthCommand) LoginCommand(bot *telego.Bot, update telego.Update) {
	//userId := tu.ID(update.CallbackQuery.From.ID)



}
