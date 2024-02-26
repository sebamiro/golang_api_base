package services

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"sync"

	"github.com/sebamiro/golang_api_base/config"
	"github.com/sebamiro/golang_api_base/pkg/funcmap"
	"github.com/sebamiro/golang_api_base/templates"
)

type (
	TemplateRenderer struct {
		templateCache sync.Map
		funcMap       template.FuncMap
		config        *config.Config
	}

	templateBuilder struct {
		build    *templateBuild
		renderer *TemplateRenderer
	}

	templateBuild struct {
		group       string
		key         string
		base        string
		files       []string
		directories []string
	}

	TemplateParsed struct {
		Template *template.Template
		build    *templateBuild
	}
)

func NewTemplateRenderer(cfg *config.Config) *TemplateRenderer {
	return &TemplateRenderer{
		templateCache: sync.Map{},
		funcMap:       funcmap.GetFuncMap(),
		config:        cfg,
	}
}

func (t *TemplateRenderer) Parse() *templateBuilder {
	log.Println("[TRACE] TemplateRenderer.Parse")
	return &templateBuilder{
		renderer: t,
		build:    &templateBuild{},
	}
}

func (t *TemplateRenderer) getCacheKey(group, key string) string {
	log.Printf("[TRACE] TemplateRenderer.getCacheKey: %s:%s\n", group, key)
	if group != "" {
		return fmt.Sprintf("%s:%s", group, key)
	}
	return key
}

func (t *TemplateRenderer) parse(build *templateBuild) (*TemplateParsed, error) {
	log.Println("[TRACE] TemplateRenderer.parse")
	var tp *TemplateParsed
	var err error

	switch {
	case build.key == "":
		return nil, errors.New("key is required")
	case len(build.files) == 0 && len(build.directories) == 0:
		return nil, errors.New("files or directories are required")
	case build.base == "":
		return nil, errors.New("base is required")
	}

	if tp, err = t.Load(build.group, build.key); err != nil || t.config.App.Environment == config.EnvLocal {
		log.Println("[TRACE] TemplateRenderer.parse: parsing template")
		parsed := template.New(build.base + config.TemplateExt).Funcs(t.funcMap)
		for k, v := range build.files {
			build.files[k] = fmt.Sprintf("%s%s", v, config.TemplateExt)
		}
		for k, v := range build.directories {
			build.directories[k] = fmt.Sprintf("%s/*%s", v, config.TemplateExt)
		}
		var tpl fs.FS
		if t.config.App.Environment == config.EnvLocal {
			tpl = templates.GetOS()
		} else {
			tpl = templates.Get()
		}
		parsed, err = parsed.ParseFS(tpl, append(build.files, build.directories...)...)
		if err != nil {
			log.Printf("[ERROR] TemplateRenderer.parse: %s\n", err)
			return nil, err
		}
		tp = &TemplateParsed{
			Template: parsed,
			build:    build,
		}
		t.templateCache.Store(t.getCacheKey(build.group, build.key), tp)
	}
	return tp, nil
}

func (t *TemplateRenderer) Load(group, key string) (*TemplateParsed, error) {
	log.Printf("[TRACE] TemplateRenderer.Load: %s:%s\n", group, key)
	load, ok := t.templateCache.Load(t.getCacheKey(group, key))
	if !ok {
		return nil, errors.New("uncached template")
	}
	templ, _ := load.(*TemplateParsed)
	return templ, nil
}

func (t *templateBuilder) Group(group string) *templateBuilder {
	log.Printf("[TRACE] templateBuilder.Group: %s\n", group)
	t.build.group = group
	return t
}

func (t *templateBuilder) Key(key string) *templateBuilder {
	log.Printf("[TRACE] templateBuilder.Key: %s\n", key)
	t.build.key = key
	return t
}

func (t *templateBuilder) Base(base string) *templateBuilder {
	log.Printf("[TRACE] templateBuilder.Base: %s\n", base)
	t.build.base = base
	return t
}

func (t *templateBuilder) Files(files ...string) *templateBuilder {
	log.Printf("[TRACE] templateBuilder.Files: %v\n", files)
	t.build.files = files
	return t
}

func (t *templateBuilder) Directories(directories ...string) *templateBuilder {
	log.Printf("[TRACE] templateBuilder.Directories: %v\n", directories)
	t.build.directories = directories
	return t
}

func (t *templateBuilder) Store() (*TemplateParsed, error) {
	log.Println("[TRACE] templateBuilder.Store")
	return t.renderer.parse(t.build)
}

func (t *templateBuilder) Execute(data any) (*bytes.Buffer, error) {
	log.Println("[TRACE] templateBuilder.Execute")
	tp, err := t.Store()
	if err != nil {
		return nil, err
	}
	return tp.Execute(data)
}

func (t *TemplateParsed) Execute(data any) (*bytes.Buffer, error) {
	log.Println("[TRACE] TemplateParsed.Execute")
	buf := new(bytes.Buffer)
	if err := t.Template.ExecuteTemplate(buf, t.build.base+config.TemplateExt, data); err != nil {
		return nil, err
	}
	return buf, nil
}
