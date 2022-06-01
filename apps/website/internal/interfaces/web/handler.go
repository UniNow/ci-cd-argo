package web

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
}

func (h Handler) ShowHome(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Worldss!")
}
