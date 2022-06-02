package web

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
}

func (h Handler) ShowHome(c echo.Context) error {
	var data interface{}
	return c.Render(http.StatusOK, "home", data)
}

func (h Handler) ShowUsers(c echo.Context) error {
	var data interface{}
	return c.Render(http.StatusOK, "users", data)
}

func (h Handler) ShowContact(c echo.Context) error {
	var data interface{}
	return c.Render(http.StatusOK, "contact", data)
}
