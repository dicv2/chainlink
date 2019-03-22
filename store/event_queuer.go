package store

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/smartcontractkit/chainlink/logger"
	"github.com/smartcontractkit/chainlink/store/models"
	"github.com/smartcontractkit/chainlink/store/orm"
)

// StatsPusher polls for events and pushes them via a WebsocketClient
type StatsPusher struct {
	ORM      *orm.ORM
	WSClient WebsocketClient
	cancel   context.CancelFunc
	Period   time.Duration
}

// NewEventQueuer returns a new event queuer
func NewEventQueuer(orm *orm.ORM, url *url.URL) *StatsPusher {
	var wsClient WebsocketClient
	wsClient = noopWebsocketClient{}
	if url != nil {
		wsClient = NewWebsocketClient(url)
	}
	return &StatsPusher{
		ORM:      orm,
		WSClient: wsClient,
		Period:   60 * time.Second,
	}
}

// Start starts the stats pusher
func (eq *StatsPusher) Start() error {
	ctx, cancel := context.WithCancel(context.Background())
	eq.cancel = cancel
	go eq.pollEvents(ctx)
	return nil
}

// Close shuts down the stats pusher
func (eq *StatsPusher) Close() error {
	if eq.cancel != nil {
		eq.cancel()
	}
	return nil
}

func (eq *StatsPusher) pollEvents(parentCtx context.Context) {
	pollTicker := time.NewTicker(eq.Period)

	for {
		select {
		case <-parentCtx.Done():
			return
		case <-pollTicker.C:
			err := eq.ORM.AllSyncEvents(func(event *models.SyncEvent) {
				fmt.Println("EventQueuer got event", event)

				eq.WSClient.Send([]byte(event.Body))

				// TODO: This is fire and forget, we may want to get confirmation
				// before deleting...

				// TODO: This should also likely have backoff logic to avoid the
				// stampeding herd problem on the link stats server

				err := eq.ORM.DB.Delete(event).Error
				if err != nil {
					logger.Errorw("Error deleting event", "event_id", event.ID, "error", err)
				}
			})

			if err != nil {
				logger.Warnf("Error querying for sync events: %v", err)
			}
		}
	}
}
