package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/testutil/assert"
)

func Test_AuthenticationMiddleware__should_abort_request_if_authentication_fails(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("GET", "/some-api", nil)
	c.Request.Header.Set("Authorization", "wrong secret")

	AuthenticationMiddleware(c)

	assert.True(t, c.IsAborted())
}

func Test_AuthenticationMiddleware__should_continue_handling_request_if_authentication_passes(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("GET", "/some-api", nil)
	c.Request.Header.Set("Authorization", "secret")

	AuthenticationMiddleware(c)

	assert.False(t, c.IsAborted())
}
