package services

import (
	"testing"

	"github.com/sebamiro/golang_api_base/config"
	"github.com/sebamiro/golang_api_base/templates"
	"github.com/stretchr/testify/assert"
)

func TestTemplateRenderer(t *testing.T) {
	group := "test"
	id := "parse"

	_, err := c.TemplateRenderer.Load(group, id)
	assert.Error(t, err)

	tpl, err := c.TemplateRenderer.
		Parse().
		Group(group).
		Key(id).
		Base("main").
		Files("layouts/main", "pages/home").
		Directories("components").
		Store()
	assert.NoError(t, err)
	parsed, err := c.TemplateRenderer.Load(group, id)
	assert.NoError(t, err)

	expectedTepmlates := make(map[string]bool)
	expectedTepmlates["main"+config.TemplateExt] = true
	expectedTepmlates["home"+config.TemplateExt] = true
	components, err := templates.Get().ReadDir("components")
	assert.NoError(t, err)
	for _, c := range components {
		expectedTepmlates[c.Name()] = true
	}
	for _, f := range parsed.Template.Templates() {
		delete(expectedTepmlates, f.Name())
	}
	assert.Empty(t, expectedTepmlates)
	data := struct {
		Name string
	}{
		Name: "Sebastian",
	}

	buf, err := tpl.Execute(data)
	assert.NoError(t, err)
	assert.NotNil(t, buf)
	assert.Contains(t, buf.String(), "Hello from "+data.Name)

	buf, err = c.TemplateRenderer.
		Parse().
		Group(group).
		Key(id).
		Base("main").
		Files("layouts/main", "pages/home").
		Directories("components").
		Execute(data)
	assert.NoError(t, err)
	assert.NotNil(t, buf)
	assert.Contains(t, buf.String(), "Hello from "+data.Name)

	_, err = c.TemplateRenderer.
		Parse().
		Store()
	assert.Error(t, err)
	_, err = c.TemplateRenderer.
		Parse().
		Execute(nil)
	assert.Error(t, err)
	_, err = c.TemplateRenderer.
		Parse().
		Key(id).
		Store()
	assert.Error(t, err)
	_, err = c.TemplateRenderer.
		Parse().
		Key(id).
		Files("layouts/main").
		Store()
	assert.Error(t, err)
	_, err = c.TemplateRenderer.
		Parse().
		Key(id).
		Directories("components").
		Store()
	assert.Error(t, err)
	_, err = c.TemplateRenderer.
		Parse().
		Key(id).
		Base("main").
		Files("AA").
		Directories("aaa").
		Store()
	assert.Error(t, err)
	_, err = c.TemplateRenderer.
		Parse().
		Key(id).
		Base("main").
		Files("AA").
		Directories("aaa").
		Execute(nil)
	assert.Error(t, err)
	_, err = c.TemplateRenderer.
		Parse().
		Key(id).
		Base("main").
		Files("layouts/main", "pages/home").
		Directories("components").
		Store()
	assert.NoError(t, err)
}
