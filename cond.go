package skyorm

// And condition aggregator.
func And(c ...Cond) Cond {
	return &cn{CondTypeAnd, nil, nil, c}
}

// Or condition aggregator.
func Or(c ...Cond) Cond {
	return &cn{CondTypeOr, nil, nil, c}
}

// Eq is equal condition.
func Eq(p Prop, value interface{}) Cond {
	return &cn{CondTypeEq, p, value, nil}
}

// Neq is not equal condition.
func Neq(p Prop, value interface{}) Cond {
	return &cn{CondTypeNeq, p, value, nil}
}

// Lt is less than condition.
func Lt(p Prop, value interface{}) Cond {
	return &cn{CondTypeLt, p, value, nil}
}

// Lte is less than or equal condition.
func Lte(p Prop, value interface{}) Cond {
	return &cn{CondTypeLte, p, value, nil}
}

// Gt is greater than condition.
func Gt(p Prop, value interface{}) Cond {
	return &cn{CondTypeGt, p, value, nil}
}

// Gte is greater than or equal condition.
func Gte(p Prop, value interface{}) Cond {
	return &cn{CondTypeGte, p, value, nil}
}

// Cond interface.
type Cond interface {
	// Type returns cond Type.
	Type() Type

	// Prop returns property definition (prop.Prop).
	Prop() Prop

	// Val returns condition value.
	Val() interface{}

	// Children returns child conditions of aggregator conditions.
	Children() []Cond
}

// Type type.
type Type uint8

const (
	CondTypeAnd Type = 1
	CondTypeOr  Type = 2
	CondTypeEq  Type = 3
	CondTypeNeq Type = 4
	CondTypeLt  Type = 5
	CondTypeLte Type = 6
	CondTypeGt  Type = 7
	CondTypeGte Type = 8
)

// cn is the default Cond implementation.
type cn struct {
	t Type
	p Prop
	v interface{}
	c []Cond
}

// Type implements Cond interface.
func (c *cn) Type() Type {
	return c.t
}

// Prop implements Cond interface.
func (c *cn) Prop() Prop {
	return c.p
}

// Val implements Cond interface.
func (c *cn) Val() interface{} {
	return c.v
}

// Children implements Cond interface.
func (c *cn) Children() []Cond {
	return c.c
}
