package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/testutil/assert"
	"zun.com/demo/models"
)

func Test_ListUserHandler__should_list_empty_users_by_default(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	ListUserHandler(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "[]", w.Body.String())
}

func Test_ListUserHandler__should_list_users(t *testing.T) {
	Users = []models.User{
		{Id: 1, Email: "mail.com", FirstName: "first"},
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	ListUserHandler(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "[{\"subject_id\":1,\"email\":\"mail.com\",\"first_name\":\"first\"}]", w.Body.String())
}
