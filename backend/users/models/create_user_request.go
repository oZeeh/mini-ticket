package models

import (
	"backend/users/enums"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserRequest struct {
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	Role     enums.Role `json:"role"`
}

func (request *CreateUserRequest) ToEntity() *User {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	return &User{
		ID:       primitive.NewObjectID(),
		Name:     request.Name,
		Email:    request.Email,
		Password: string(hashed),
		Role:     request.Role,
	}
}
