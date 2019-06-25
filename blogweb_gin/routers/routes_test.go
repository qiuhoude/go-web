package routers

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInitRouter(t *testing.T) {
	router := InitRouter()

	rw := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/register", nil)

	//http.DefaultServeMux.ServeHTTP(rw,req) // 原生的http

	router.ServeHTTP(rw, req)

	assert.Equal(t, 200, rw.Code)
	t.Log(rw.Body.String())
	assert.Equal(t, "pong", rw.Body.String())
}
