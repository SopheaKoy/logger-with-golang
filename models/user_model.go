package models

import (
	"time"

	"github.com/google/uuid"
)

// UserModel represents a user in the system
type UserModel struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	FirstName   string    `json:"first_name" gorm:"not null"`
	LastName    string    `json:"last_name" gorm:"not null"`
	Email       string    `json:"email" gorm:"unique;not null"`
	Password	string    `json:"-" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt	time.Time `json:"deleted_at" gorm:"autoDeleteTime"`
}

/*
In Go, return &UserModel{} is used to return a pointer to an instance of the UserModel struct. Here's what it does:

Breakdown:
	&UserModel{}:
	&: The & operator in Go is used to get the memory address of a variable. When used with a struct (in this case, UserModel), it returns a pointer to that struct.
	{}: The curly braces {} are used to initialize the struct. The values inside the curly braces are assigned to the fields of the struct. If no values are provided, 
the struct fields are set to their zero values (e.g., empty strings for strings, 0 for integers, nil for pointers, etc.).
return:
	The return keyword is used to return a value from a function. In this case, return &UserModel{} returns a pointer to the newly created UserModel instance.
*/

func UserCreation(firstName, lastName, email, Password string) *UserModel {
	return &UserModel{
		ID			: uuid.New(),
		FirstName	: firstName,
		LastName	: lastName,
		Email		: email,
		Password	: Password,
	}
}