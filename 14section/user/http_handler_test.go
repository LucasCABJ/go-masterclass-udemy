package user

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	Handler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	body, err := io.ReadAll(rr.Body)
	assert.NoError(t, err)
	assert.Equal(t, "Hello World", string(body))
}

func TestEchoHandler(t *testing.T) {
	rr := httptest.NewRecorder()
	reqBody := bytes.NewBufferString("Hello Echo!")
	req := httptest.NewRequest("GET", "/", reqBody)

	HandleEcho(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	body, err := io.ReadAll(rr.Body)
	assert.NoError(t, err)
	assert.Equal(t, "Hello Echo!", string(body))
}
