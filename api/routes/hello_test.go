package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type hello struct {
	Message string `json:"message"`
}

func TestRun(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/hello", Hello)

	req, _ := http.NewRequest("GET", "/hello", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusOK)

	expected := hello{
		Message: "The answer to life, the universe and everything is 42.",
	}

	actual := hello{}
	json.Unmarshal([]byte(resp.Body.String()), &actual)

	assert.Equal(t, expected, actual)
}
