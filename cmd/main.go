//
// smiro
// 2024-02
//

package main

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"os/signal"

	"github.com/sebamiro/Gym-FP/pkg/services"

	"github.com/labstack/echo/v4"
)

type Template struct {
	template *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.template.ExecuteTemplate(w, name, data)
}

func main() {
	c, err := services.NewContainer()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := c.Shutdown(); err != nil {
			c.Web.Logger.Fatal(err)
		}
	}()

	t := &Template{
		template: template.Must(template.ParseGlob("templates/*.html")),
	}
	c.Web.Renderer = t
	c.Web.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{})
	})

	go func() {
		srv := http.Server{
			Addr:         fmt.Sprintf("%s:%d", c.Config.HTTP.Hostname, c.Config.HTTP.Port),
			Handler:      c.Web,
			ReadTimeout:  c.Config.HTTP.ReadTimeout,
			WriteTimeout: c.Config.HTTP.WrtieTimeout,
			IdleTimeout:  c.Config.HTTP.IdleTimeout,
		}
		// TLS
		if err := c.Web.StartServer(&srv); err != nil {
			c.Web.Logger.Fatalf("shutting down the server with error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, os.Kill)
	<-quit
	if err := c.Web.Shutdown(context.Background()); err != nil {
		c.Web.Logger.Fatal(err)
	}
}
