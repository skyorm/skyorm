package skyorm

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		name string
		typ  string
		isPk bool
	}
	tests := []struct {
		name string
		args args
		want Prop
	}{
		{
			name: "prop1",
			args: args{"prop1", "string", false},
			want: &pr{"prop1", "string", false},
		},
		{
			name: "prop2",
			args: args{"prop2", "int", true},
			want: &pr{"prop2", "int", true},
		},
		{
			name: "prop3",
			args: args{"prop3", "uint64", true},
			want: &pr{"prop3", "uint64", true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProp(tt.args.name, tt.args.typ, tt.args.isPk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_property_IsPk(t *testing.T) {
	type fields struct {
		name string
		typ  string
		isPk bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "true",
			fields: fields{"prop1", "string", true},
			want:   true,
		},
		{
			name:   "false",
			fields: fields{"prop1", "string", true},
			want:   true,
		},
		{
			name:   "nil",
			fields: fields{},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pr{
				name: tt.fields.name,
				typ:  tt.fields.typ,
				isPk: tt.fields.isPk,
			}
			if got := p.IsPk(); got != tt.want {
				t.Errorf("IsPk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_property_Name(t *testing.T) {
	type fields struct {
		name string
		typ  string
		isPk bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "prop1",
			fields: fields{"prop1", "string", true},
			want:   "prop1",
		},
		{
			name:   "prop2",
			fields: fields{"prop2", "string", true},
			want:   "prop2",
		},
		{
			name:   "nil",
			fields: fields{},
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pr{
				name: tt.fields.name,
				typ:  tt.fields.typ,
				isPk: tt.fields.isPk,
			}
			if got := p.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_property_Type(t *testing.T) {
	type fields struct {
		name string
		typ  string
		isPk bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "prop1",
			fields: fields{"prop1", "string", false},
			want:   "string",
		},
		{
			name:   "prop2",
			fields: fields{"prop2", "int", true},
			want:   "int",
		},
		{
			name:   "prop3",
			fields: fields{"prop3", "uint64", true},
			want:   "uint64",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pr{
				name: tt.fields.name,
				typ:  tt.fields.typ,
				isPk: tt.fields.isPk,
			}
			if got := p.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}
