package pinterest

import (
	"errors"
	"github.com/carrot/go-pinterest"
)

type PinterestServiceApi struct {
	Client *pinterest.Client
}

func NewPinterestService(accessToken string) (*PinterestServiceApi, error) {
	if accessToken == "" {
		return nil, errors.New("token is empty")
	}

	newClient := pinterest.
		NewClient().
		RegisterAccessToken(accessToken)

	return &PinterestServiceApi{
		Client: newClient,
	}, nil
}

// Generate access token
func (p *PinterestServiceApi) Authenticate(
	clientId string,
	clientSecret string,
	accessCode string,
) (*PinterestServiceApi, error) {

	_, err := p.Client.OAuth.Token.Create(
		clientId,
		clientSecret,
		accessCode,
	)
	if err != nil {
		return nil, err
	}

	return p, nil
}
