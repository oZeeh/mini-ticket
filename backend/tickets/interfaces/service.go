package interfaces

import (
	"backend/tickets/models"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	Create(ctx context.Context, u *models.CreateTicketRequest) (primitive.ObjectID, error)
	FindAll(ctx context.Context) ([]models.TicketEntity, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (*models.TicketEntity, error)
	FindByAssignedTechnitian(ctx context.Context, id primitive.ObjectID) ([]models.TicketEntity, error)
	FindByUser(ctx context.Context, id primitive.ObjectID) ([]models.TicketEntity, error)
	FindDoneTickets(ctx context.Context, id primitive.ObjectID) ([]models.TicketEntity, error)
	FindOpenTickets(ctx context.Context, id primitive.ObjectID) ([]models.TicketEntity, error)
	Assign(ctx context.Context, ticketID primitive.ObjectID, technicianID primitive.ObjectID) error
	Update(ctx context.Context, u *models.UpdateTicketRequest, userID primitive.ObjectID) error
	Delete(ctx context.Context, ticketID primitive.ObjectID, userID primitive.ObjectID) error
}
