package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/testutil/assert"
	controllers "zun.com/demo/controllers"
	"zun.com/demo/models"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func StringPtr(v string) *string { return &v }

func Test_ListUserHandler__should_return_empty_list(t *testing.T) {
	mockResponse := `[]`
	r := SetUpRouter()
	r.GET("/users", controllers.ListUserHandler)

	req, _ := http.NewRequest("GET", "/users", nil)
	req.Header.Add("Authorization", "secret")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func Test_ListUserHandler__should_list_an_item_after_added(t *testing.T) {
	r := SetUpRouter()
	r.GET("/users", controllers.ListUserHandler)
	r.POST("/users", controllers.AddUserHandler)

	userDto := models.UserDto{
		Email:     "zun@zun.com",
		FirstName: "abc",
		LastName:  StringPtr("he"),
		Address:   StringPtr("ho"),
	}

	userDtoJson, _ := json.Marshal(userDto)
	postReq, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(userDtoJson))
	postReq.Header.Add("Authorization", "secret")
	w1 := httptest.NewRecorder()
	r.ServeHTTP(w1, postReq)

	var newUserDto models.NewUserDto
	json.Unmarshal(w1.Body.Bytes(), &newUserDto)
	assert.NotNil(t, newUserDto.Id)
	assert.IsKind(t, reflect.Int, newUserDto.Id)
	assert.Equal(t, http.StatusCreated, w1.Code)

	getReq, _ := http.NewRequest("GET", "/users", nil)
	getReq.Header.Add("Authorization", "secret")
	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, getReq)

	var users []models.User
	json.Unmarshal(w2.Body.Bytes(), &users)
	assert.NotNil(t, users)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, userDto.FirstName, users[0].FirstName)
	assert.Equal(t, userDto.LastName, users[0].LastName)
	assert.Equal(t, userDto.Email, users[0].Email)
	assert.Equal(t, userDto.Address, users[0].Address)
}

// TODO: more tests for other scenarios here
