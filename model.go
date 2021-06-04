package skyorm

// Model interface.
type Model interface {
	OrmStore() Store
	OrmPk() interface{}
	OrmPkProp() Prop
	OrmPkPointer() interface{}
	OrmProps() []Prop
	OrmPointers() []interface{}
	OrmVals() []interface{}
}
