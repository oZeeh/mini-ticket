package models

import (
	"backend/tickets/enums"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateTicketRequest struct {
	Title       string              `json:"title" binding:"required"`
	Description string              `json:"description" binding:"required"`
	UserID      primitive.ObjectID  `json:"-"`
	AssignedTo  *primitive.ObjectID `json:"-"`
	Status      enums.TicketStatus  `json:"-"`
}

func (request *CreateTicketRequest) ToEntity() *TicketEntity {
	return &TicketEntity{
		Title:       request.Title,
		Description: request.Description,
		UserID:      request.UserID,
		AssignedTo:  nil,
		Status:      enums.Open,
	}
}
