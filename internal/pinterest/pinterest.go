package pinterest

import (
	"errors"
	"github.com/carrot/go-pinterest"
)

type PinterestServiceApi struct {
	Client *pinterest.Client
}

func NewPinterestService(
	clientId string,
	clientSecret string,
	token string,
) (*PinterestServiceApi, error) {
	newClient := pinterest.NewClient()

	if token == "" {
		return nil, errors.New("token is empty")
	}

	_, err := newClient.OAuth.Token.Create(
		clientId,
		clientSecret,
		token,
	)
	if err != nil {
		return nil, err
	}

	newClient.RegisterAccessToken(token)

	return &PinterestServiceApi{
		Client: newClient,
	}, nil
}
