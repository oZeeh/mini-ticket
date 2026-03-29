package interfaces

import (
	"context"

	"backend/users/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	Create(ctx context.Context, u *models.User) (primitive.ObjectID, error)
	FindAll(ctx context.Context) ([]models.User, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	Update(ctx context.Context, u *models.User) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}
