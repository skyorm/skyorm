package skyorm

import (
	`context`
	`errors`
)

// NewDB returns a new DB for Provider.
func NewDB(provider Provider) DB {
	return &db{provider}
}

// DB interface.
type DB interface {
	// Put Model(s) into the database.
	Put(models ...Model) error

	// PutContext puts Model(s) into the database with context.
	PutContext(ctx context.Context, models ...Model) error

	// Populate a Model by primary key.
	Populate(model Model, pk interface{}) error

	// PopulateContext populates a Model by primary key with context.
	PopulateContext(ctx context.Context, model Model, pk interface{}) error

	// Find Model(s) in database by condition and limit.
	Find(store Store, condition Cond, limit, offset int) ([]Model, error)

	// Find finds Model(s) in database by condition and limit with context.
	FindContext(ctx context.Context, store Store, condition Cond, limit, offset int) ([]Model, error)

	// Update the Store by condition.
	Update(store Store, condition Cond, values ...Val) error

	// UpdateContext updates a Store by condition with context.
	UpdateContext(ctx context.Context, store Store, condition Cond, values ...Val) error

	// Delete from Store by condition.
	Delete(store Store, condition Cond) error

	// DeleteContext deletes from Store by condition with context.
	DeleteContext(ctx context.Context, store Store, condition Cond) error

	// Count records in Store by condition.
	Count(store Store, condition Cond) (int64, error)

	// CountContext counts records in Store by condition with context.
	CountContext(ctx context.Context, store Store, condition Cond) (int64, error)

	// ErrNotFound returns provider's NotFound error.
	ErrNotFound() error

	// IsNotFound checks if the error equals to provider's NotFound error.
	IsNotFound(err error) bool
}

type db struct {
	p Provider
}

func (d *db) Put(models ...Model) error {
	return d.PutContext(context.Background(), models...)
}

func (d *db) PutContext(ctx context.Context, models ...Model) error {
	if len(models) == 0 {
		return errors.New("at least one model is required in Put method")
	}
	return d.p.Put(ctx, models...)
}

func (d *db) Populate(model Model, pk interface{}) error {
	return d.PopulateContext(context.Background(), model, pk)
}

func (d *db) PopulateContext(ctx context.Context, model Model, pk interface{}) error {
	return d.p.Populate(ctx, model, pk)
}

func (d *db) Find(store Store, condition Cond, limit, offset int) ([]Model, error) {
	return d.FindContext(context.Background(), store, condition, limit, offset)
}

func (d *db) FindContext(ctx context.Context, store Store, condition Cond, limit, offset int) ([]Model, error) {
	return d.p.Find(ctx, store, condition, limit, offset)
}

func (d *db) Update(store Store, condition Cond, values ...Val) error {
	return d.UpdateContext(context.Background(), store, condition, values...)
}

func (d *db) UpdateContext(ctx context.Context, store Store, condition Cond, values ...Val) error {
	return d.p.Update(ctx, store, condition, values...)
}

func (d *db) Delete(store Store, condition Cond) error {
	return d.DeleteContext(context.Background(), store, condition)
}

func (d *db) DeleteContext(ctx context.Context, store Store, condition Cond) error {
	return d.p.Delete(ctx, store, condition)
}

func (d *db) Count(store Store, condition Cond) (int64, error) {
	return d.CountContext(context.Background(), store, condition)
}

func (d *db) CountContext(ctx context.Context, store Store, condition Cond) (int64, error) {
	return d.p.Count(ctx, store, condition)
}

func (d *db) ErrNotFound() error {
	return d.p.ErrNotFound()
}

func (d *db) IsNotFound(err error) bool {
	return d.ErrNotFound() == err
}
