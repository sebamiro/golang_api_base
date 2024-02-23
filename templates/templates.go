package templates

import "embed"

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

var templates embed.FS

func Get() embed.FS {
	return templates
}
