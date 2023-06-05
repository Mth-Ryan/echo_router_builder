package router_builder

import (
	"strings"

	"github.com/labstack/echo/v4"
)

type RouterBuilder struct {
	inner *echo.Echo
}

func NewBuilder(middlewares ...echo.MiddlewareFunc) *RouterBuilder {
	inner := echo.New()
	inner.Use(middlewares...)

	return &RouterBuilder{inner}
}

func parseRoute(r string) string {
	if r == "/" || r == "" {
		return ""
	} else if !strings.HasPrefix(r, "/") {
		return "/" + r
	}
	return r
}

func (b *RouterBuilder) Register(c *Controller) *RouterBuilder {
	g := b.inner.Group(parseRoute(c.baseRoute), c.middlewares...)

	for _, h := range c.handlers {
		g.Add(h.method, parseRoute(h.route), h.inner, h.middlewares...)
	}

	return b
}

func (b *RouterBuilder) RegisterStatic(route, folder string) *RouterBuilder {
	b.inner.Static(parseRoute(route), folder)
	return b
}

func (b *RouterBuilder) Build() *echo.Echo {
	return b.inner
}
