package routes

import (
	"github.com/sebamiro/golang_api_base/pkg/controller"
	"github.com/sebamiro/golang_api_base/templates"

	"github.com/labstack/echo/v4"
)

const routeNameHome = "home"
type home struct {
	controller.Controller
}

func (c *home) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = templates.LayoutMain
	page.Name = templates.PageHome
	return c.RenderPage(ctx, page)
}
