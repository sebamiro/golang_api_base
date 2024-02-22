//
// smiro
// 2024-02
//

package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Template struct {
	template *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.template.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	t := &Template{
		template: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = t
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{})
	})
	e.Logger.Fatal(e.Start(":8080"))
}
