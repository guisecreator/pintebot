package types

import (
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegoapi"
	"github.com/mymmrac/telego/telegoapi/mock"
)

type HandlersSessionManager[T Session] struct {
	mock.MockCaller
}

func (s *HandlersSessionManager[T]) Get() (*telegoapi.Response, error) {
	response, err := s.MockCaller.
		Call("Get",
			&telegoapi.RequestData{
				ContentType: "Content-Type",
				Buffer:      nil,
			})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *HandlersSessionManager[T]) Remove() error {
	panic("implement me")
	return nil
}

func (s *HandlersSessionManager[T]) Update(update *telego.Update) error {
	panic("implement me")
	return nil
}
