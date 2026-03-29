package models

import (
	"backend/tickets/enums"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TicketEntity struct {
	ID          primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	Title       string              `bson:"title" json:"title"`
	Description string              `bson:"description" json:"description"`
	UserID      primitive.ObjectID  `bson:"user_id" json:"user_id"`
	AssignedTo  *primitive.ObjectID `bson:"assigned_to,omitempty" json:"assigned_to,omitempty"`
	Status      enums.TicketStatus  `bson:"status" json:"status"`
}
