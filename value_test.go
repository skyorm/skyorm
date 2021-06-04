package skyorm

import (
	"reflect"
	"testing"
)

func TestNewVal(t *testing.T) {
	p1 := NewProp("s", "string", false)
	type args struct {
		prop Prop
		v    interface{}
	}
	tests := []struct {
		name string
		args args
		want Val
	}{
		{
			name: "nil",
			args: args{p1, nil},
			want: &vl{p1, nil},
		},
		{
			name: "string",
			args: args{p1, "string"},
			want: &vl{p1, "string"},
		},
		{
			name: "int",
			args: args{p1, 1},
			want: &vl{p1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVal(tt.args.prop, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_vl_Prop(t *testing.T) {
	p1 := NewProp("s", "string", false)
	p2 := NewProp("i", "int", false)
	type fields struct {
		p Prop
		v interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   Prop
	}{

		{
			name:   "nil",
			fields: fields{},
			want:   nil,
		},
		{
			name:   "string",
			fields: fields{p1, "string"},
			want:   &pr{"s", "string", false},
		},
		{
			name:   "string",
			fields: fields{p2, "string"},
			want:   &pr{"i", "int", false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &vl{
				p: tt.fields.p,
				v: tt.fields.v,
			}
			if got := v.Prop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_vl_Val(t *testing.T) {
	p1 := NewProp("s", "string", false)
	type fields struct {
		p Prop
		v interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name:   "nil",
			fields: fields{p1, nil},
			want:   nil,
		},
		{
			name:   "string",
			fields: fields{p1, "string"},
			want:   "string",
		},
		{
			name:   "int",
			fields: fields{p1, 1},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &vl{
				p: tt.fields.p,
				v: tt.fields.v,
			}
			if got := v.Val(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Val() = %v, want %v", got, tt.want)
			}
		})
	}
}
