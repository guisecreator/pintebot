package middlewares

import th "github.com/mymmrac/telego/telegohandler"

type ChatMiddleware struct {
	Next th.Handler
}

//TODO: create chat middleware
//func (m *ChatMiddleware) Wrap(next th.Handler) th.Handler {
//	return func(ctx context.Context, update string) error {
//		return
//	}
//}
