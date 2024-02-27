package controller

import (
	"net/http"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/sebamiro/golang_api_base/config"
	"github.com/sebamiro/golang_api_base/pkg/services"
	"github.com/sebamiro/golang_api_base/pkg/tests"
	"github.com/stretchr/testify/assert"
)

var c *services.Container

func TestMain(m *testing.M) {
	var err error
	config.SwitchEnvironment(config.EnvTest)
	c, err = services.NewContainer()
	if err != nil {
		panic(err)
	}
	code := m.Run()
	if err := c.Shutdown(); err != nil {
		panic(err)
	}
	os.Exit(code)
}

func TestController_Redirect(t *testing.T) {
	c.Web.GET("/path/:first/:second", func(ctx echo.Context) error {
		return nil
	}).Name = "redirect-test"

	ctx, _ := tests.NewContext(c.Web, "/abc")
	ctrl := NewController(c)
	err := ctrl.Redirect(ctx, "redirect-test", "one", "two")
	assert.NoError(t, err)
	assert.Equal(t, "/path/one/two", ctx.Response().Header().Get(echo.HeaderLocation))
	assert.Equal(t, http.StatusFound, ctx.Response().Status)
}
