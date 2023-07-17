package mongodb

import (
	"context"
	"fmt"

	"github.com/PDeXchange/pac/internal/pkg/pac-go-server/models"
	"go.mongodb.org/mongo-driver/bson"
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
func (db *MongoDB) GetEventsByUserID(id string) ([]models.Event, error) {
	events := []models.Event{}

	filter := bson.D{{}}
	if id != "" {
		filter = bson.D{{Key: "user_id", Value: id}}
	}

	collection := db.Database.Collection("events")
	ctx, cancel := context.WithTimeout(context.Background(), dbContextTimeout)
	defer cancel()
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error getting events: %w", err)
	}
	defer cur.Close(ctx)

	if err = cur.All(context.Background(), &events); err != nil {
		return nil, fmt.Errorf("error fetching events: %w", err)
	}

	return events, nil
}
