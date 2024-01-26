package handlers

import (
	"github.com/guisecreator/pintebot/internal/telegram/types"
	"log"
)

type UserPinsCommand struct {
	*types.CommandsOptions
}

func (u *UserPinsCommand) GetAccountPins() {
	account_pins, err := u.Services.PinterestClient.GetPinByid("")
	if err != nil {
		log.Println(err)
	}

	log.Println(account_pins)

}

//https://www.pinterest.com/oauth/?&client_id=1478464&redirect_uri=https://pinshuffle.fly.dev/redirect/&response_type=code&scope=user_accounts:read,catalogs:read,boards:read,boards:read_secret,pins:read,pins:read_secret
