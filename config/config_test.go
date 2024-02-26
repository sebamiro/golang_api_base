package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	_, err := GetConfig()
	assert.NoError(t, err)

	var env environment
	env = "xxx"
	SwitchEnvironment(env)
	cfg, err := GetConfig()
	assert.NoError(t, err)
	assert.Equal(t, env, cfg.App.Environment)
}
