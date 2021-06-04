package skyorm

import (
	"context"
)

// Provider interface.
type Provider interface {
	// Put is adding one or more records into the Provider, reproduced by Model.
	Put(ctx context.Context, models ...Model) error

	// Populate is population property values from the Provider into the Model,
	// identified by pk (primary key).
	Populate(ctx context.Context, model Model, pk interface{}) error

	// Find is searching the Provider for data, filtered by Cond. Optionally,
	// a limit and an offset can be set. Offset will be taken into consideration,
	// only if limit is > 0.
	Find(ctx context.Context, store Store, condition Cond, limit, offset int) ([]Model, error)

	// Update Val-s in the Provider, filtered by Cond.
	Update(ctx context.Context, store Store, condition Cond, values ...Val) error

	// Delete records in the Provider, filtered by Cond.
	Delete(ctx context.Context, store Store, condition Cond) error

	// Count the amount of records in the Provider, filtered by Cond.
	Count(ctx context.Context, store Store, condition Cond) (int64, error)

	// ErrNotFound returns the Provider's "not found" error.
	ErrNotFound() error
}
