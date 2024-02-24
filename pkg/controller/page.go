package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sebamiro/Gym-FP/templates"
)

type Page struct {
	AppName string
	Title   string
	Context echo.Context
	ToURL   func(name string, params ...any) string
	Path    string
	URL     string
	Data    any
	Layout  templates.Layout
	Name    templates.Page
	// Form    any
	// TODO USER
	// IsAuth bool
	// AuthUser any
	StatusCode int
	Metatags   struct {
		Description string
		Keywords    []string
	}
}

func NewPage(ctx echo.Context) Page {
	return Page{
		Context:    ctx,
		ToURL:      ctx.Echo().Reverse,
		Path:       ctx.Request().URL.Path,
		URL:        ctx.Request().URL.String(),
		StatusCode: http.StatusOK,
	}
}
