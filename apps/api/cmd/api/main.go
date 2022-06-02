package main

import (
	"github.com/labstack/echo/v4"
	"main/internal/interfaces/rest"
)

func main() {
	e := echo.New()
	rest.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
