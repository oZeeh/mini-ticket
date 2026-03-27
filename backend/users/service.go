package users

import (
	"backend/users/interfaces"
	"backend/users/models"
	"context"
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

	println("to aqui")
	user := request.RequestToEntity()

	println(user)
	return s.repo.Create(ctx, user)
}

func (s *UserService) Find(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return s.repo.FindByID(ctx, id)
}

func (s *UserService) Update(ctx context.Context, request *models.CreateUserRequest) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user := request.RequestToEntity()

	err := s.repo.Update(ctx, user)

	return user, err
}

func (s *UserService) Delete(ctx context.Context, id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return s.repo.Delete(ctx, id)
}
