package job

import (
	"context"
	"net/url"

	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

//go:generate mockery --quiet --name ServiceCtx --output ./mocks/ --case=underscore

type Service interface {
	Start() error
	Close() error
}

// ServiceCtx is the same as Service, but Start method receives a context.
type ServiceCtx interface {
	// Start starts the service.
	// Start should not block; any long-running operations should be started in a goroutine.
	// The context passed to Start should be used to cancel the startup process.
	// Do not use the passed context object after the Start method returns.
	Start(context.Context) error

	// Close stops the service and releases any resources it holds.
	// Close should block until the goroutines spawned by the service returned.
	// The system will call Close only if the Start method completed without an error and will never attempt to restart a service after it has been closed.
	Close() error
}

type Config interface {
	URL() url.URL
	pg.QConfig
}

// ServiceAdapter is a helper introduced for transitioning from Service to ServiceCtx.
type ServiceAdapter interface {
	ServiceCtx
}

type adapter struct {
	service Service
}

// NewServiceAdapter creates an adapter instance for the given Service.
func NewServiceAdapter(service Service) ServiceCtx {
	return &adapter{
		service,
	}
}

// Start forwards the call to the underlying service.Start().
// Context is not used in this case.
func (a adapter) Start(context.Context) error {
	return a.service.Start()
}

// Close forwards the call to the underlying service.Close().
func (a adapter) Close() error {
	return a.service.Close()
}
