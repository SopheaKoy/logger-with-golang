package schema

import (
	model "logger/models"

	"github.com/google/uuid"
)

type UserCreation struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func UserSchemaOut(firstName, lastName, email, Password string) *model.UserModel {
	return &model.UserModel{
		ID			: uuid.New(),
		FirstName	: firstName,
		LastName	: lastName,
		Email		: email,
		Password	: Password,
	}
}