package monitoring

import (
	"context"
	"sync"
)

type FeedMonitor interface {
	Run(ctx context.Context)
}

func NewFeedMonitor(
	log Logger,
	pollers []Poller,
	exporters []Exporter,
) FeedMonitor {
	return &feedMonitor{
		log,
		pollers,
		exporters,
	}
}

type feedMonitor struct {
	log       Logger
	pollers   []Poller
	exporters []Exporter
}

// Run should be executed as a goroutine.
// Signal termination by cancelling ctx; then wait for Run() to exit.
func (f *feedMonitor) Run(ctx context.Context) {
	f.log.Infow("starting feed monitor")
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	updates := make(chan interface{})
	wg.Add(len(f.pollers))
	for _, poller := range f.pollers {
		go func(poller Poller) {
			defer wg.Done()
			select {
			case update := <-poller.Updates():
				select {
				case updates <- update:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}(poller)
	}

	// Consume updates.
	for {
		var update interface{}
		select {
		case update = <-updates:
		case <-ctx.Done():
			f.cleanup()
			return
		}
		// TODO (dru) do we need a worker pool here?
		wg.Add(len(f.exporters))
		for _, exp := range f.exporters {
			go func(exp Exporter) {
				defer wg.Done()
				exp.Export(ctx, update)
			}(exp)
		}
	}
}

func (f *feedMonitor) cleanup() {
	for _, exp := range f.exporters {
		exp.Cleanup()
	}
}
