package skyorm

// NewStore returns new Store.
func NewStore(name string, pkIndex int, factory ModelFactory, props ...Prop) Store {
	return &st{name, pkIndex, props, factory}
}

// Store interface.
type Store interface {
	Name() string
	Props() []Prop
	Pk() Prop
	Model() Model
}

// ModelFactory type.
type ModelFactory func() Model

// st is the default Store implementation.
type st struct {
	n  string
	pk int
	pr []Prop
	f  ModelFactory
}

// Name implements Store interface.
func (s *st) Name() string {
	return s.n
}

// Props implements Store interface.
func (s *st) Props() []Prop {
	return s.pr
}

// Pk implements Store interface.
func (s *st) Pk() Prop {
	return s.pr[s.pk]
}

// Model implements Store interface.
func (s *st) Model() Model {
	return s.f()
}
