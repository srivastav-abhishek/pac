package main

import (
	"go.uber.org/zap"

	"github.com/PDeXchange/pac/internal/pkg/notifier/client"
	"github.com/PDeXchange/pac/internal/pkg/pac-go-server/db"
	log "github.com/PDeXchange/pac/internal/pkg/pac-go-server/logger"
	"github.com/PDeXchange/pac/internal/pkg/pac-go-server/models"
)

func notifier(db db.DB, mailClient client.Notifier) {
	l := log.GetLogger()
	eventCh := make(chan *models.Event)
	go func() {
		err := db.WatchEvents(eventCh)
		if err != nil {
			l.Fatal("error watching events", zap.Error(err))
		}
	}()

	for event := range eventCh {
		l.Info("Received an event:", zap.Any("Document", event))
		// skip if already notified
		if event.Notified {
			l.Info("Already notified")
			continue
		}
		l.Info("Notifying")
		if err := mailClient.Notify(*event); err != nil {
			l.Error("Error notifying", zap.Error(err))
		}
		if err := db.MarkEventAsNotified(event.ID.Hex()); err != nil {
			l.Error("Error updating event", zap.Error(err))
		}
	}
}
