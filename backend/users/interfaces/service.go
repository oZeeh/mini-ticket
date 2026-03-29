package interfaces

import (
	"backend/users/models"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	Create(ctx context.Context, request *models.CreateUserRequest) (primitive.ObjectID, error)
	Find(ctx context.Context, id string) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	Update(ctx context.Context, request *models.CreateUserRequest) (*models.User, error)
	Delete(ctx context.Context, id string) error
}
