package tickets

import (
	customErrors "backend/custom_errors"
	"backend/tickets/interfaces"
	"backend/tickets/models"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ticketService struct {
	repo interfaces.Repository
}

func NewService(repo interfaces.Repository) interfaces.Service {
	return &ticketService{repo: repo}
}

func (s *ticketService) Create(ctx context.Context, u *models.CreateTicketRequest) (primitive.ObjectID, error) {
	entity := u.ToEntity()
	return s.repo.Create(ctx, entity)
}

func (s *ticketService) FindAll(ctx context.Context) ([]models.TicketEntity, error) {
	return s.repo.FindAll(ctx)
}

func (s *ticketService) FindByID(ctx context.Context, id primitive.ObjectID) (*models.TicketEntity, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *ticketService) FindByAssignedTechnitian(ctx context.Context, id primitive.ObjectID) ([]models.TicketEntity, error) {
	return s.repo.FindByAssignedTechnitian(ctx, id)
}

func (s *ticketService) FindByUser(ctx context.Context, id primitive.ObjectID) ([]models.TicketEntity, error) {
	return s.repo.FindByUser(ctx, id)
}

func (s *ticketService) FindDoneTickets(ctx context.Context, id primitive.ObjectID) ([]models.TicketEntity, error) {
	return s.repo.FindDoneTickets(ctx, id)
}

func (s *ticketService) FindOpenTickets(ctx context.Context, id primitive.ObjectID) ([]models.TicketEntity, error) {
	return s.repo.FindOpenTickets(ctx, id)
}

func (s *ticketService) Update(ctx context.Context, u *models.UpdateTicketRequest, id primitive.ObjectID) error {
	entity := u.ToEntity(id)
	return s.repo.Update(ctx, entity)
}

func (s *ticketService) Assign(ctx context.Context, ticketID primitive.ObjectID, technicianID primitive.ObjectID) error {
	ticket, err := s.repo.FindByID(ctx, ticketID)
	if err != nil {
		return err
	}
	if ticket == nil {
		return customErrors.ErrNotFound
	}
	if ticket.AssignedTo != nil {
		return customErrors.ErrConflict // já tem técnico assignado
	}
	return s.repo.Assign(ctx, ticketID, technicianID)
}

func (s *ticketService) Delete(ctx context.Context, ticketID primitive.ObjectID, userID primitive.ObjectID) error {
	ticket, err := s.repo.FindByID(ctx, ticketID)
	if err != nil {
		return err
	}
	if ticket == nil {
		return customErrors.ErrNotFound
	}
	if ticket.UserID != userID {
		return customErrors.ErrForbidden
	}
	return s.repo.Delete(ctx, ticketID)
}
