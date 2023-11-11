package pinterest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPinterestService_Success(t *testing.T) {
	clientId := "your_client_id"
	clientSecret := "your_client_secret"
	token := "your_token"

	service, err := NewPinterestService(clientId, clientSecret, token)

	assert.NoError(t, err)
	assert.NotNil(t, service)
	assert.NotNil(t, service.Client)
}

func TestNewPinterestService_EmptyTokenError(t *testing.T) {
	clientId := "your_client_id"
	clientSecret := "your_client_secret"
	token := ""

	service, err := NewPinterestService(clientId, clientSecret, token)

	assert.Error(t, err)
	assert.Nil(t, service)
	assert.EqualError(t, err, "token is empty")
}
