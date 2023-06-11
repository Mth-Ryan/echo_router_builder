package router_builder

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	method      string
	route       string
	inner       echo.HandlerFunc
	middlewares []echo.MiddlewareFunc
}

func newHandler(method, route string, inner echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) Handler {
	return Handler{method, route, inner, middlewares}
}

type Controller struct {
	baseRoute   string
	handlers    []Handler
	middlewares []echo.MiddlewareFunc
}

func NewController(baseRoute string, middlewares ...echo.MiddlewareFunc) *Controller {
	return &Controller{
		baseRoute:   baseRoute,
		middlewares: middlewares,
		handlers:    []Handler{},
	}
}

func (c *Controller) addHandler(method, route string, inner echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) {
	handler := newHandler(method, route, inner, middlewares...)
	c.handlers = append(c.handlers, handler)
}

func (c *Controller) Get(route string, inner echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) *Controller {
	c.addHandler(http.MethodGet, route, inner, middlewares...)
	return c
}

func (c *Controller) Post(route string, inner echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) *Controller {
	c.addHandler(http.MethodPost, route, inner, middlewares...)
	return c
}

func (c *Controller) Put(route string, inner echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) *Controller {
	c.addHandler(http.MethodPut, route, inner, middlewares...)
	return c
}

func (c *Controller) Patch(route string, inner echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) *Controller {
	c.addHandler(http.MethodPatch, route, inner, middlewares...)
	return c
}

func (c *Controller) Delete(route string, inner echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) *Controller {
	c.addHandler(http.MethodDelete, route, inner, middlewares...)
	return c
}

func (c *Controller) Head(route string, inner echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) *Controller {
	c.addHandler(http.MethodHead, route, inner, middlewares...)
	return c
}

func (c *Controller) Options(route string, inner echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) *Controller {
	c.addHandler(http.MethodOptions, route, inner, middlewares...)
	return c
}

func (c *Controller) View(route string, view string, args any, middlewares ...echo.MiddlewareFunc) *Controller {
	c.Get(route, func(c echo.Context) error {
		return c.Render(http.StatusOK, view, args)
	}, middlewares...)

	return c
}

func (c *Controller) ViewWitchCB(route string, view string, dataCB func(echo.Context) any, middlewares ...echo.MiddlewareFunc) *Controller {
	c.Get(route, func(c echo.Context) error {
		return c.Render(http.StatusOK, view, dataCB(c))
	}, middlewares...)

	return c
}

func (c *Controller) RedirectGet(route string, to string, status int) *Controller {
	c.Get(route, func(c echo.Context) error {
		return c.Redirect(status, to)
	})
	return c
}

func (c *Controller) RedirectPost(route string, to string, status int) *Controller {
	c.Post(route, func(c echo.Context) error {
		return c.Redirect(status, to)
	})
	return c
}

func (c *Controller) RedirectPut(route string, to string, status int) *Controller {
	c.Put(route, func(c echo.Context) error {
		return c.Redirect(status, to)
	})
	return c
}

func (c *Controller) RedirectDelete(route string, to string, status int) *Controller {
	c.Delete(route, func(c echo.Context) error {
		return c.Redirect(status, to)
	})
	return c
}

func (c *Controller) RedirectPatch(route string, to string, status int) *Controller {
	c.Patch(route, func(c echo.Context) error {
		return c.Redirect(status, to)
	})
	return c
}
