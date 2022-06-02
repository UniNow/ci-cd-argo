package web

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterRoutes(e *echo.Echo) {
	h := Handler{}
	e.GET("/", h.ShowHome)
	e.GET("/users", h.ShowUsers)
	e.GET("/contactsss", h.ShowContact)
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Worldssss!")
	})
}
