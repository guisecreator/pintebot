package pinterest

import (
	"github.com/guisecreator/plient"
	"github.com/guisecreator/plient/models"
)

type PinterestClient struct {
	Client plient.Client
}

func NewPinterestClient(AccessToken, ClientId, ClientSecret string) (*PinterestClient, error) {
	client, err := plient.NewClient(
		AccessToken,
		ClientId,
		ClientSecret,
	)
	if err != nil {
		return nil, err
	}

	// err = client.Authorize()
	// if err != nil {
	// 	return nil, err
	// }

	return &PinterestClient{
		Client: *client,
	}, nil
}

func (c *PinterestClient) GetPinByid(id string) (*models.PinsData, error) {
	pin, err := c.Client.GetPinById(id)
	if err != nil {
		return nil, err
	}

	return pin, nil
}
