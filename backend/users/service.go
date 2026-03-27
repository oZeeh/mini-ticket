package users

import (
	"backend/users/interfaces"
	"backend/users/models"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	repo interfaces.Repository
}

func NewService(r interfaces.Repository) *UserService {
	return &UserService{
		repo: r,
	}
}

func (s *UserService) Create(ctx context.Context, request *models.CreateUserRequest) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user := request.RequestToEntity()

	return s.repo.Create(ctx, user)
}

func (s *UserService) Find(ctx context.Context, id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}

	return s.repo.FindByID(ctx, objectID)
}

func (s *UserService) Update(ctx context.Context, request *models.CreateUserRequest) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user := request.RequestToEntity()

	err := s.repo.Update(ctx, user)

	return user, err
}

func (s *UserService) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id")
	}

	return s.repo.Delete(ctx, objectID)
}
