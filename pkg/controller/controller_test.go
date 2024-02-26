package controller

import (
	"os"
	"testing"

	"github.com/sebamiro/golang_api_base/config"
	"github.com/sebamiro/golang_api_base/pkg/services"
)

var c *services.Container

func TestMain(m *testing.M) {

	config.SwitchEnvironment(config.EnvTest)
	c, err := services.NewContainer()
	if err != nil {
		panic(err)
	}
	code := m.Run()
	if err := c.Shutdown(); err != nil {
		panic(err)
	}
	os.Exit(code)
}
