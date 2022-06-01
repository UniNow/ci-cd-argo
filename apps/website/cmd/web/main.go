package main

import (
	"github.com/labstack/echo/v4"
	"main/internal/interfaces/web"
)

func main() {
	e := echo.New()
	e.Logger.Error(web.RegisterRoutes(e))
	e.Logger.Fatal(e.Start(":1323"))
}
