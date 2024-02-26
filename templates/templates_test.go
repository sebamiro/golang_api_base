package templates

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	_, err := Get().Open(fmt.Sprintf("pages/%s.html", PageHome))
	assert.NoError(t, err)
}

func TestGetOS(t *testing.T) {
	_, err := GetOS().Open(fmt.Sprintf("pages/%s.html", PageHome))
	assert.NoError(t, err)
}

