package recovery

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRecoveryMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	engine := gin.New()
	engine.Use(Middleware())
	engine.GET("/panic", func(context *gin.Context) {
		panic("Unexpected panic!")
	})

	httpRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/panic", nil)
	assert.NoError(t, err)

	assert.NotPanics(t, func() { engine.ServeHTTP(httpRecorder, req) })
}
