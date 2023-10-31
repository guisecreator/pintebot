package types

import (
	"context"
	"github.com/celestix/gotgproto/dispatcher/handlers/filters"
	"github.com/mymmrac/telego"
)

type SessionManager[T comparable] interface {
	Get(ctx context.Context) *T
	Update(update telego.Update) telego.Update
	Filter() filters.InlineQueryFilter
	Remove(ctx context.Context) error
}

type Session struct {
	Token string
}
