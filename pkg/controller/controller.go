package controller

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sebamiro/Gym-FP/pkg/services"
)

type Controller struct {
	Container *services.Container
}

func NewController(c *services.Container) Controller {
	return Controller{Container: c}
}

func (c *Controller) RenderPage(ctx echo.Context, page Page) error {
	var buf *bytes.Buffer
	var err error
	templateGroup := "page"

	if page.Name == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "Page name is required")
	}
	if page.AppName == "" {
		page.AppName = c.Container.Config.App.Name
	}
	buf, err = c.Container.TempalteRenderer.
		Parse().
		Group(templateGroup).
		Key(string(page.Name)).
		Base(string(page.Layout)).
		Files(
			fmt.Sprintf("layouts/%s", string(page.Layout)),
			fmt.Sprintf("pages/%s", string(page.Name)),
		).
		Directories("components").
		Execute(page)
	if err != nil {
		return c.Fail(err, "failed to render page")
	}
	ctx.Response().Status = page.StatusCode
	return ctx.HTMLBlob(ctx.Response().Status, buf.Bytes())
}

func (c *Controller) Redirect(ctx echo.Context, route string, routeParams ...any) error {
	url := ctx.Echo().Reverse(route, routeParams...)
	return ctx.Redirect(http.StatusFound, url)
}

func (c *Controller) Fail(err error, log string) error {
	return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%s: %v", log, err))
}
