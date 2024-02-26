package routes

import (
	"net/http"

	"github.com/sebamiro/golang_api_base/pkg/controller"
	"github.com/sebamiro/golang_api_base/pkg/services"

	echow "github.com/labstack/echo/v4/middleware"
)

func BuildRouter(c *services.Container) {
	g := c.Web.Group("")
	g.Use(
		echow.RemoveTrailingSlashWithConfig(echow.TrailingSlashConfig{
			RedirectCode: http.StatusMovedPermanently,
		}),
		echow.Recover(),
		echow.Secure(),
		echow.RequestID(),
		echow.Gzip(),
		echow.Logger(),
		echow.TimeoutWithConfig(echow.TimeoutConfig{
			Timeout: c.Config.App.Timeout,
		}),
	)
	ctr := controller.NewController(c)

	// Error handler
	// Routes
	home := home{ctr}
	g.GET("/", home.Get).Name = routeNameHome
}
