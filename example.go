package example

import (
	"github.com/goflash/flash/v2"
	"github.com/goflash/flash/v2/ctx"
)

// ExampleMiddleware is an example of a middleware function.
func ExampleMiddleware(next flash.Handler) flash.Handler {
	return func(c flash.Ctx) error {
		l := ctx.LoggerFromContext(c.Context())
		l.Info("ExampleMiddleware: before next")
		return next(c)
	}
}
