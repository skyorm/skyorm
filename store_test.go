package skyorm

import (
	"reflect"
	"testing"
)

func TestNewStore(t *testing.T) {
	p := NewProp("s", "string", false)
	type args struct {
		name       string
		pkIndex    int
		factory    ModelFactory
		properties []Prop
	}
	tests := []struct {
		name string
		args args
		want Store
	}{
		{
			name: "empty",
			args: args{
				name: "empty",
			},
			want: &st{
				n: "empty",
			},
		},
		{
			name: "withPk",
			args: args{
				name:    "withPk",
				pkIndex: 1,
			},
			want: &st{
				n:  "withPk",
				pk: 1,
			},
		},
		{
			name: "withProps",
			args: args{
				name:       "withProps",
				pkIndex:    1,
				properties: []Prop{p},
			},
			want: &st{
				n:  "withProps",
				pk: 1,
				pr: []Prop{p},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStore(tt.args.name, tt.args.pkIndex, tt.args.factory, tt.args.properties...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_store_Model(t *testing.T) {
	fNil := func() Model {
		return nil
	}
	fModel := func() Model {
		return new(demoModel)
	}
	type fields struct {
		name       string
		pkIndex    int
		properties []Prop
		factory    ModelFactory
	}
	tests := []struct {
		name   string
		fields fields
		want   Model
	}{
		{
			name:   "empty",
			fields: fields{demoStore.Name(), 0, demoStore.Props(), fNil},
			want:   nil,
		},
		{
			name:   "model",
			fields: fields{demoStore.Name(), 0, demoStore.Props(), fModel},
			want:   &demoModel{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &st{
				n:  tt.fields.name,
				pk: tt.fields.pkIndex,
				pr: tt.fields.properties,
				f:  tt.fields.factory,
			}
			if got := s.Model(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Model() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_store_Name(t *testing.T) {
	fNil := func() Model {
		return nil
	}
	type fields struct {
		name       string
		pkIndex    int
		properties []Prop
		factory    ModelFactory
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "demo",
			fields: fields{demoStore.Name(), 0, demoStore.Props(), fNil},
			want:   "demo",
		},
		{
			name:   "demo",
			fields: fields{"other", 0, demoStore.Props(), fNil},
			want:   "other",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &st{
				n:  tt.fields.name,
				pk: tt.fields.pkIndex,
				pr: tt.fields.properties,
				f:  tt.fields.factory,
			}
			if got := s.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_store_Pk(t *testing.T) {
	props := []Prop{
		NewProp("one", "string", false),
		NewProp("two", "int", false),
		NewProp("three", "uint64", false),
	}
	type fields struct {
		pkIndex    int
		properties []Prop
	}
	tests := []struct {
		name   string
		fields fields
		want   Prop
	}{
		{
			name: "one",
			fields: fields{
				pkIndex:    0,
				properties: props,
			},
			want: props[0],
		},
		{
			name: "two",
			fields: fields{
				pkIndex:    1,
				properties: props,
			},
			want: props[1],
		},
		{
			name: "three",
			fields: fields{
				pkIndex:    2,
				properties: props,
			},
			want: props[2],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &st{
				pk: tt.fields.pkIndex,
				pr: tt.fields.properties,
			}
			if got := s.Pk(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_store_Props(t *testing.T) {

	tests := []struct {
		name  string
		props []Prop
		want  []Prop
	}{
		{
			name:  "empty",
			props: []Prop{},
			want:  []Prop{},
		},
		{
			name: "one",
			props: []Prop{
				NewProp("one", "string", false),
			},
			want: []Prop{
				NewProp("one", "string", false),
			},
		},
		{
			name: "one",
			props: []Prop{
				NewProp("one", "string", false),
				NewProp("two", "int", false),
			},
			want: []Prop{
				NewProp("one", "string", false),
				NewProp("two", "int", false),
			},
		},
		{
			name: "one",
			props: []Prop{
				NewProp("one", "string", false),
				NewProp("two", "int", false),
				NewProp("three", "uint64", false),
			},
			want: []Prop{
				NewProp("one", "string", false),
				NewProp("two", "int", false),
				NewProp("three", "uint64", false),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &st{
				pr: tt.props,
			}
			if got := s.Props(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Props() = %v, want %v", got, tt.want)
			}
		})
	}
}

var (
	demoStore = NewStore("demo", 0,
		func() Model {
			return new(demoModel)
		},
		NewProp("id", "int", true),
		NewProp("v", "string", false))
)
