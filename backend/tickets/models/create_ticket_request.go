package models

import (
	"backend/tickets/enums"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateTicketRequest struct {
	Title       string              `json:"title"`
	Description string              `json:"description"`
	UserID      primitive.ObjectID  `json:"user_id"`
	AssignedTo  *primitive.ObjectID `json:"assigned_to"`
	Status      enums.TicketStatus  `json:"status"`
}

func (request *CreateTicketRequest) ToEntity() *TicketEntity {
	return &TicketEntity{
		Title:       request.Title,
		Description: request.Description,
		UserID:      request.UserID,
		AssignedTo:  nil,
		Status:      request.Status,
	}
}
