package middlewares

import (
	th "github.com/mymmrac/telego/telegohandler"
)

type Middleware struct {
	Next th.Handler
}
