package main

import (
	"github.com/Masterminds/sprig/v3"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"main/internal/interfaces/web"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Debug = true
	e.Static("/assets", "public/assets")
	t := &Template{
		templates: template.Must(
			template.New("layout").Funcs(sprig.FuncMap()).ParseGlob("views/*.gohtml"),
		),
	}
	e.Renderer = t
	web.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
