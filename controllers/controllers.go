package controllers

import (
	"math"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	slices "golang.org/x/exp/slices"
	m "zun.com/demo/models"
)

// TODO: use real database here
var Users = []m.User{}

func AddUserHandler(c *gin.Context) {
	var userDto m.UserDto
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	index := slices.IndexFunc(Users, func(user m.User) bool {
		return user.Email == userDto.Email
	})

	if index != -1 {
		c.JSON(http.StatusConflict, gin.H{
			"error": "email is in use",
		})
		return
	}

	newUser := m.User{
		Id:        rand.Intn(math.MaxInt32),
		Email:     userDto.Email,
		FirstName: userDto.FirstName,
		LastName:  userDto.LastName,
		Address:   userDto.Address,
	}
	Users = append(Users, newUser)

	c.JSON(http.StatusCreated, m.NewUserDto{
		Id: newUser.Id,
	})
}

func UpdateUserHandler(c *gin.Context) {
	var userDto m.UserDto
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	index := slices.IndexFunc(Users, func(user m.User) bool {
		return user.Email == userDto.Email
	})

	if index == -1 {
		c.JSON(http.StatusConflict, gin.H{
			"error": "user does not exist",
		})
		return
	}

	Users[index] = m.User{
		Id:        Users[index].Id,
		Email:     userDto.Email,
		FirstName: userDto.FirstName,
		LastName:  userDto.LastName,
		Address:   userDto.Address,
	}

	c.JSON(http.StatusNoContent, nil)
}

func ListUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, Users)
}
