package skyorm

// NewProp returns a new Prop.
func NewProp(name, typ string, isPk bool) Prop {
	return &pr{name, typ, isPk}
}

// Prop interface.
type Prop interface {
	Name() string
	Type() string
	IsPk() bool
}

type pr struct {
	name string
	typ  string
	isPk bool
}

// Name implements Prop interface.
func (p *pr) Name() string {
	return p.name
}

// Type implements Prop interface.
func (p *pr) Type() string {
	return p.typ
}

// IsPk implements Prop interface.
func (p *pr) IsPk() bool {
	return p.isPk
}
