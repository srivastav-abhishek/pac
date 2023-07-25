package client

import "github.com/PDeXchange/pac/internal/pkg/pac-go-server/models"

type Notifier interface {
	Notify(event models.Event) error
}
