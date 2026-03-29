package users

import (
	"backend/users/interfaces"
	"backend/users/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) interfaces.Repository {
	return &MongoRepository{
		collection: db.Collection("users"),
	}
}

// FindByEmail implements [interfaces.Repository].
func (r *MongoRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user models.User
	result := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)

	if result != nil {
		return nil, result
	}

	return &user, nil
}

// Create implements [interfaces.Repository].
func (r *MongoRepository) Create(ctx context.Context, u *models.User) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := r.collection.InsertOne(ctx, u)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

// FindAll implements [interfaces.Repository].
func (r *MongoRepository) FindAll(ctx context.Context) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var users []models.User

	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

// FindByID implements [interfaces.Repository].
func (r *MongoRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user models.User
	err := r.collection.FindOne(
		ctx,
		bson.M{"_id": id},
	).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	return &user, err
}

// Update implements [interfaces.Repository].
func (r *MongoRepository) Update(ctx context.Context, u *models.User) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.collection.UpdateByID(
		ctx,
		u.ID,
		bson.M{
			"$set": bson.M{
				"name":  u.Name,
				"email": u.Email,
			},
		},
	)

	return err
}

// Delete implements [interfaces.Repository].
func (r *MongoRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(
		ctx,
		bson.M{"_id": id},
	)

	return err
}
