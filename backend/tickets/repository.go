package tickets

import (
	"backend/tickets/enums"
	"backend/tickets/interfaces"
	"backend/tickets/models"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) interfaces.Repository {
	return &mongoRepository{
		collection: db.Collection("tickets"),
	}
}

func (r *mongoRepository) Create(ctx context.Context, u *models.TicketEntity) (primitive.ObjectID, error) {
	u.ID = primitive.NewObjectID()
	_, err := r.collection.InsertOne(ctx, u)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return u.ID, nil
}

func (r *mongoRepository) FindAll(ctx context.Context) ([]models.TicketEntity, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tickets []models.TicketEntity
	if err := cursor.All(ctx, &tickets); err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *mongoRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.TicketEntity, error) {
	var ticket models.TicketEntity
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&ticket)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}
	return &ticket, err
}

func (r *mongoRepository) FindByAssignedTechnitian(ctx context.Context, id primitive.ObjectID) ([]models.TicketEntity, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"assigned_to": id})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tickets []models.TicketEntity
	if err := cursor.All(ctx, &tickets); err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *mongoRepository) FindByUser(ctx context.Context, id primitive.ObjectID) ([]models.TicketEntity, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"user_id": id})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tickets []models.TicketEntity
	if err := cursor.All(ctx, &tickets); err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *mongoRepository) FindDoneTickets(ctx context.Context, id primitive.ObjectID) ([]models.TicketEntity, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"userid": id, "status": enums.Done})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tickets []models.TicketEntity
	if err := cursor.All(ctx, &tickets); err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *mongoRepository) FindOpenTickets(ctx context.Context, id primitive.ObjectID) ([]models.TicketEntity, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"userid": id, "status": enums.Open})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tickets []models.TicketEntity
	if err := cursor.All(ctx, &tickets); err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *mongoRepository) Update(ctx context.Context, u *models.TicketEntity) error {
	_, err := r.collection.ReplaceOne(ctx, bson.M{"_id": u.ID}, u)
	return err
}

func (r *mongoRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *mongoRepository) Assign(ctx context.Context, ticketID primitive.ObjectID, technicianID primitive.ObjectID) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": ticketID},
		bson.M{"$set": bson.M{"assigned_to": technicianID}},
	)
	return err
}
