package main

import (
	"go.uber.org/zap"

	"github.com/PDeXchange/pac/internal/pkg/notifier/client/mail"
	"github.com/PDeXchange/pac/internal/pkg/pac-go-server/db/mongodb"
	log "github.com/PDeXchange/pac/internal/pkg/pac-go-server/logger"
)

func main() {
	l := log.GetLogger()
	l.Info("Starting event notifier")
	db := mongodb.New()
	if err := db.Connect(); err != nil {
		l.Fatal("Error connecting to MongoDB", zap.Error(err))
	}
	defer db.Disconnect()
	mailClient := mail.New()
	notifier(db, mailClient)
}
