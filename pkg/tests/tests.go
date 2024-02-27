package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/labstack/echo/v4"
)

// NewContext creates a new Echo context for tests using an HTTP test request and response recorder
func NewContext(e *echo.Echo, url string) (echo.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(http.MethodGet, url, strings.NewReader(""))
	rec := httptest.NewRecorder()
	return e.NewContext(req, rec), rec
}
