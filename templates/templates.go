package templates

import (
	"embed"
	"runtime"
	"path"
	"os"
	"path/filepath"
	"io/fs"
)

type (
	Layout string
	Page   string
)

const (
	LayoutMain Layout = "main"
	// LayoutAuth Layout = "auth"
	// LayoutHTMX Layout = "htmx"
)

const (
	PageHome  Page = "home"
	PageLogin Page = "login"
)

//go:embed *
var templates embed.FS

func Get() embed.FS {
	return templates
}

// Local environment
func GetOS() fs.FS {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	p := filepath.Join(filepath.Dir(d), "templates")
	return os.DirFS(p)
}
