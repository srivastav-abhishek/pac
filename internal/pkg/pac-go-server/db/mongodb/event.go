package mongodb

import (
	"context"
	"fmt"

	"github.com/PDeXchange/pac/internal/pkg/pac-go-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db *MongoDB) NewEvent(e *models.Event) error {
	collection := db.Database.Collection("events")
	ctx, cancel := context.WithTimeout(context.Background(), dbContextTimeout)
	defer cancel()
	_, err := collection.InsertOne(ctx, e)
	if err != nil {
		return fmt.Errorf("error inserting Event: %w", err)
	}

	return nil
}

// GetEventsByUserID returns all events by user id
func (db *MongoDB) GetEventsByUserID(id string, startIndex, perPage int64) ([]models.Event, int64, error) {
	events := []models.Event{}
	var totalCount int64

	filter := bson.D{{}}
	if id != "" {
		filter = bson.D{{Key: "user_id", Value: id}}
	}

	findOptions := options.Find()
	findOptions.SetSkip(startIndex)
	findOptions.SetLimit(perPage)

	collection := db.Database.Collection("events")
	ctx, cancel := context.WithTimeout(context.Background(), dbContextTimeout)
	defer cancel()
	cur, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, totalCount, fmt.Errorf("error getting events: %w", err)
	}
	defer cur.Close(ctx)

	if err = cur.All(context.Background(), &events); err != nil {
		return nil, totalCount, fmt.Errorf("error fetching events: %w", err)
	}
	// Get the total number of events from the database
	totalCount, err = collection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		return nil, totalCount, fmt.Errorf("error getting total count of events: %w", err)
	}

	return events, totalCount, nil
}
