package models

type User struct {
	Id        int     `json:"subject_id"`
	Email     string  `json:"email"`
	FirstName string  `json:"first_name"`
	LastName  *string `json:"last_name,omitempty"`
	Address   *string `json:"address,omitempty"`
}

type UserDto struct {
	Email     string  `json:"email" binding:"required"`
	FirstName string  `json:"first_name" binding:"required"`
	LastName  *string `json:"last_name"`
	Address   *string `json:"address"`
}

type NewUserDto struct {
	Id int `json:"subject_id"`
}
