// auth/interfaces/service.go
package interfaces

import (
	"backend/auth/models"
	"context"
)

type Service interface {
	Login(ctx context.Context, req *models.LoginRequest) (*models.LoginResponse, error)
}
