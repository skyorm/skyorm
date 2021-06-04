package skyorm

type demoModel struct {
	ID int
	V  string
}

func (d *demoModel) OrmStore() Store {
	return demoStore
}

func (d *demoModel) OrmPk() interface{} {
	return d.ID
}

func (d *demoModel) OrmPkProp() Prop {
	return demoStore.Pk()
}

func (d *demoModel) OrmPkPointer() interface{} {
	return &d.ID
}

func (d *demoModel) OrmProps() []Prop {
	return demoStore.Props()
}

func (d *demoModel) OrmPointers() []interface{} {
	return []interface{}{
		&d.ID,
		&d.V,
	}
}

func (d *demoModel) OrmVals() []interface{} {
	return []interface{}{
		d.ID,
		d.V,
	}
}
