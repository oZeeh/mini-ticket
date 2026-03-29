package models

import (
	"backend/tickets/enums"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateTicketRequest struct {
	ID          primitive.ObjectID  `json:"id"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	AssignedTo  *primitive.ObjectID `json:"assigned_to"`
	Status      enums.TicketStatus  `json:"status"`
}

func (request *UpdateTicketRequest) ToEntity(userID primitive.ObjectID) *TicketEntity {
	return &TicketEntity{
		ID:          request.ID,
		Title:       request.Title,
		Description: request.Description,
		UserID:      userID,
		AssignedTo:  request.AssignedTo,
		Status:      request.Status,
	}
}
