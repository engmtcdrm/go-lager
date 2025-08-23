package lager

import (
	"context"
	"log/slog"
)

type Router struct {
	handlers []slog.Handler
}

func NewRouter() *Router {
	return &Router{
		handlers: []slog.Handler{},
	}
}

func (r *Router) Add(h slog.Handler, match func(ctx context.Context, r slog.Record) bool) {
	r.handlers = append(r.handlers, h)
}
