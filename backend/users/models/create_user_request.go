package models

import "backend/users/enums"

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     enums.Role
}

func (request *CreateUserRequest) RequestToEntity() *User {
	return &User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Role:     request.Role,
	}
}
