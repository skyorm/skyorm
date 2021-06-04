package skyorm

// NewVal returns new Val.
func NewVal(prop Prop, v interface{}) Val {
	return &vl{prop, v}
}

// Val interface.
type Val interface {
	// Prop returns Prop.
	Prop() Prop

	// Val returns vl.
	Val() interface{}
}

type vl struct {
	p Prop
	v interface{}
}

// Prop implements Val interface.
func (v *vl) Prop() Prop {
	return v.p
}

// Val implements Val interface.
func (v *vl) Val() interface{} {
	return v.v
}
