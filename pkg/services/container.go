package services

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/sebamiro/Gym-FP/config"
)

type Container struct {
	Config *config.Config
	Web    *echo.Echo
	// DB
	// Template
}

func NewContainer() (*Container, error) {
	c := new(Container)
	if err := c.initConfig(); err != nil {
		return nil, err
	}
	c.Web = echo.New()
	return c, nil
}

func (c *Container) initConfig() error {
	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}
	log.Printf("[TRACE] initConfig: %v\n", cfg)
	c.Config = &cfg
	return nil
}

func (c *Container) Shutdown() error {
	// c.DB.Close()
	log.Println("[TRACE] Shutdown")
	return nil
}
