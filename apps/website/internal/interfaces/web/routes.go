package web

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) error {
	h := Handler{}
	e.GET("/", h.ShowHome)
	e.GET("/health", h.ShowHome)
	return nil
}
