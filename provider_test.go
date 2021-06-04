package skyorm

import (
	"context"
	"errors"
	"testing"
)

var (
	demoProviderErr = errors.New("demoProviderErrored")
)

type demoProviderOK struct {
	t *testing.T
}

func (p *demoProviderOK) Put(ctx context.Context, models ...Model) error {
	//if ctx == nil {
	//	p.t.Error("ctx is nil")
	//}
	//if len(models) == 0 {
	//	p.t.Error("models list empty")
	//}
	//if !reflect.DeepEqual(models, []Model{demoStore.Model()}) {
	//	p.t.Error("not a demo model")
	//}
	return nil
}

func (p *demoProviderOK) Populate(ctx context.Context, model Model, pk interface{}) error {
	//if ctx == nil {
	//	p.t.Error("ctx is nil")
	//}
	//if !reflect.DeepEqual(model, []Model{demoStore.Model()}) {
	//	p.t.Error("not a demo model")
	//}
	//if reflect.ValueOf(pk).IsNil() {
	//	p.t.Error("pk is nil")
	//}
	return nil
}

func (p *demoProviderOK) Find(ctx context.Context, store Store, condition Cond, limit, offset int) ([]Model, error) {
	return make([]Model, 0), nil
}

func (p *demoProviderOK) Update(ctx context.Context, store Store, condition Cond, values ...Val) error {
	return nil
}

func (p *demoProviderOK) Delete(ctx context.Context, store Store, condition Cond) error {
	return nil
}

func (p *demoProviderOK) Count(ctx context.Context, store Store, condition Cond) (int64, error) {
	return 1, nil
}

func (p *demoProviderOK) ErrNotFound() error {
	return demoProviderErr
}

type demoProviderErrored struct {
	t *testing.T
}

func (p *demoProviderErrored) Put(ctx context.Context, models ...Model) error {
	return errors.New("error")
}

func (p *demoProviderErrored) Populate(ctx context.Context, model Model, pk interface{}) error {
	return errors.New("error")
}

func (p *demoProviderErrored) Find(ctx context.Context, store Store, condition Cond, limit, offset int) ([]Model, error) {
	return nil, errors.New("error")
}

func (p *demoProviderErrored) Update(ctx context.Context, store Store, condition Cond, values ...Val) error {
	return errors.New("error")
}

func (p *demoProviderErrored) Delete(ctx context.Context, store Store, condition Cond) error {
	return errors.New("error")
}

func (p *demoProviderErrored) Count(ctx context.Context, store Store, condition Cond) (int64, error) {
	return 0, errors.New("error")
}

func (p *demoProviderErrored) ErrNotFound() error {
	return errors.New("error")
}
