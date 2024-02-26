package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewContainer(t *testing.T) {
	assert.NotNil(t, c)
	assert.NotNil(t, c.Web)
	assert.NotNil(t, c.Config)
	assert.NotNil(t, c.TemplateRenderer)
}
