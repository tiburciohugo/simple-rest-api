package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHelloUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/hello/*name", HelloUser)

	req, _ := http.NewRequest("GET", "/hello/John", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.JSONEq(t, `{"message":"Hello, John!"}`, resp.Body.String())

	req, _ = http.NewRequest("GET", "/hello/", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 400, resp.Code)
	assert.JSONEq(t, `{"message":"name is required"}`, resp.Body.String())
}
