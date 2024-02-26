package services

import (
	"testing"
	"os"

	"github.com/sebamiro/golang_api_base/config"
)

var (
	c *Container
)

func TestMain(m *testing.M) {
	var err error
	config.SwitchEnvironment(config.EnvTest)
	c, err = NewContainer()
	if err != nil {
		panic(err)
	}
	code := m.Run()
	if err = c.Shutdown(); err != nil {
		panic(err)
	}
	os.Exit(code)
}
