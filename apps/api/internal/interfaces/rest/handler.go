package rest

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type User struct {
	ID int `json:"id"`
}

type UserCollection struct {
	Items []User
}

type JsonResponse struct {
	Data interface{} `json:"data"`
}

type Handler struct {
	UserCollection UserCollection
}

func (h Handler) UserIndex(c echo.Context) error {
	return c.JSON(http.StatusOK, JsonResponse{Data: h.UserCollection.Items})
}
