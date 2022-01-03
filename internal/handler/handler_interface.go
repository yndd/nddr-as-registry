package handler

import (
	"context"

	"github.com/yndd/ndd-runtime/pkg/logging"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Option can be used to manipulate Options.
type Option func(Handler)

// WithLogger specifies how the Reconciler should log messages.
func WithLogger(log logging.Logger) Option {
	return func(s Handler) {
		s.WithLogger(log)
	}
}

func WithClient(c client.Client) Option {
	return func(s Handler) {
		s.WithClient(c)
	}
}

type Handler interface {
	WithLogger(log logging.Logger)
	WithClient(a client.Client)
	Init(string, uint32, uint32, string)
	Delete(string)
	GetAllocated(string) (uint32, []*uint32)
	ResetSpeedy(string)
	GetSpeedy(crName string) int
	IncrementSpeedy(crName string)
	Register(context.Context, *RegisterInfo) (*uint32, error)
	DeRegister(context.Context, *RegisterInfo) error
}
