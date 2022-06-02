package rest

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterRoutes(e *echo.Echo) {
	h := Handler{UserCollection: UserCollection{Items: []User{}}}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Worldss!")
	})
	e.GET("/api/users", h.UserIndex)

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Health!")
	})
}
