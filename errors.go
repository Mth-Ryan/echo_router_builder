package router_builder

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ErrorHandler struct {
	handlers map[int]echo.HTTPErrorHandler
}

func NewErrorHandler() *ErrorHandler {
	defaultHandlers := map[int]echo.HTTPErrorHandler{
		http.StatusNotFound: func(err error, c echo.Context) {
			c.String(http.StatusNotFound, "Not found")
		},
		http.StatusUnauthorized: func(err error, c echo.Context) {
			c.String(http.StatusUnauthorized, "Unauthorized")
		},
		http.StatusInternalServerError: func(err error, c echo.Context) {
			c.String(http.StatusInternalServerError, "Internal server error")
		},
	}

	return &ErrorHandler{handlers: defaultHandlers}
}

func (e *ErrorHandler) AddHandler(status int, handler echo.HTTPErrorHandler) {
	e.handlers[status] = handler
}

func (e *ErrorHandler) GetHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		he, ok := err.(*echo.HTTPError)
		if !ok {
			he = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		code := he.Code
		message := he.Message.(string)

		if h, ok := e.handlers[code]; ok {
			h(err, c)
		} else {
			c.String(code, message)
		}
	}
}
