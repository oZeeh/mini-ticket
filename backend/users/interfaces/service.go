package interfaces

import (
	"backend/users/models"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	Create(ctx context.Context, request *models.CreateUserRequest) (primitive.ObjectID, error)
	Find(ctx context.Context, id primitive.ObjectID) (*models.User, error)
	Update(ctx context.Context, request *models.CreateUserRequest) (*models.User, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
}
