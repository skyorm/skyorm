package skyorm

import (
	"reflect"
	"testing"
)

func TestAnd(t *testing.T) {
	p := NewProp("s", "string", false)
	type args struct {
		c []Cond
	}
	tests := []struct {
		name string
		args args
		want Cond
	}{
		{
			name: "empty",
			want: &cn{t: CondTypeAnd},
		},
		{
			name: "with children",
			args: args{[]Cond{
				Eq(p, "1"),
				Neq(p, "0"),
			}},
			want: &cn{
				t: CondTypeAnd,
				c: []Cond{
					&cn{t: CondTypeEq, p: p, v: "1"},
					&cn{t: CondTypeNeq, p: p, v: "0"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := And(tt.args.c...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("And() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEq(t *testing.T) {
	ps := NewProp("s", "string", false)
	pi := NewProp("i", "int", false)
	pj := NewProp("i", "json.RawMessage", false)
	type args struct {
		p     Prop
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want Cond
	}{
		{
			"string",
			args{ps, "string"},
			&cn{t: CondTypeEq, p: ps, v: "string"},
		},
		{
			"int",
			args{pi, "int"},
			&cn{t: CondTypeEq, p: pi, v: "int"},
		},
		{
			"json",
			args{pj, "json"},
			&cn{t: CondTypeEq, p: pj, v: "json"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Eq(tt.args.p, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Eq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGt(t *testing.T) {
	ps := NewProp("s", "string", false)
	pi := NewProp("i", "int", false)
	pj := NewProp("i", "json.RawMessage", false)
	type args struct {
		p     Prop
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want Cond
	}{
		{
			"string",
			args{ps, "string"},
			&cn{t: CondTypeGt, p: ps, v: "string"},
		},
		{
			"int",
			args{pi, "int"},
			&cn{t: CondTypeGt, p: pi, v: "int"},
		},
		{
			"json",
			args{pj, "json"},
			&cn{t: CondTypeGt, p: pj, v: "json"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Gt(tt.args.p, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Gt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGte(t *testing.T) {
	ps := NewProp("s", "string", false)
	pi := NewProp("i", "int", false)
	pj := NewProp("i", "json.RawMessage", false)
	type args struct {
		p     Prop
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want Cond
	}{
		{
			"string",
			args{ps, "string"},
			&cn{t: CondTypeGte, p: ps, v: "string"},
		},
		{
			"int",
			args{pi, "int"},
			&cn{t: CondTypeGte, p: pi, v: "int"},
		},
		{
			"json",
			args{pj, "json"},
			&cn{t: CondTypeGte, p: pj, v: "json"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Gte(tt.args.p, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Gte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLt(t *testing.T) {
	ps := NewProp("s", "string", false)
	pi := NewProp("i", "int", false)
	pj := NewProp("i", "json.RawMessage", false)
	type args struct {
		p     Prop
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want Cond
	}{
		{
			"string",
			args{ps, "string"},
			&cn{t: CondTypeLt, p: ps, v: "string"},
		},
		{
			"int",
			args{pi, "int"},
			&cn{t: CondTypeLt, p: pi, v: "int"},
		},
		{
			"json",
			args{pj, "json"},
			&cn{t: CondTypeLt, p: pj, v: "json"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lt(tt.args.p, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLte(t *testing.T) {
	ps := NewProp("s", "string", false)
	pi := NewProp("i", "int", false)
	pj := NewProp("i", "json.RawMessage", false)
	type args struct {
		p     Prop
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want Cond
	}{
		{
			"string",
			args{ps, "string"},
			&cn{t: CondTypeLte, p: ps, v: "string"},
		},
		{
			"int",
			args{pi, "int"},
			&cn{t: CondTypeLte, p: pi, v: "int"},
		},
		{
			"json",
			args{pj, "json"},
			&cn{t: CondTypeLte, p: pj, v: "json"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lte(tt.args.p, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNeq(t *testing.T) {
	ps := NewProp("s", "string", false)
	pi := NewProp("i", "int", false)
	pj := NewProp("i", "json.RawMessage", false)
	type args struct {
		p     Prop
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want Cond
	}{
		{
			"string",
			args{ps, "string"},
			&cn{t: CondTypeNeq, p: ps, v: "string"},
		},
		{
			"int",
			args{pi, "int"},
			&cn{t: CondTypeNeq, p: pi, v: "int"},
		},
		{
			"json",
			args{pj, "json"},
			&cn{t: CondTypeNeq, p: pj, v: "json"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Neq(tt.args.p, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Neq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOr(t *testing.T) {
	p := NewProp("s", "string", false)
	type args struct {
		c []Cond
	}
	tests := []struct {
		name string
		args args
		want Cond
	}{
		{
			name: "empty",
			want: &cn{t: CondTypeOr},
		},
		{
			name: "with children",
			args: args{[]Cond{
				Eq(p, "1"),
				Neq(p, "0"),
			}},
			want: &cn{
				t: CondTypeOr,
				c: []Cond{
					&cn{t: CondTypeEq, p: p, v: "1"},
					&cn{t: CondTypeNeq, p: p, v: "0"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Or(tt.args.c...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Or() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cn_Children(t *testing.T) {
	tests := []struct {
		name     string
		children []Cond
		want     []Cond
	}{
		{"empty", nil, nil},
		{"emptyStruct", []Cond{&cn{}}, []Cond{&cn{}}},
		{"emptyStruct", []Cond{&cn{t: CondTypeEq}}, []Cond{&cn{t: CondTypeEq}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cn{
				c: tt.children,
			}
			if got := c.Children(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Children() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cn_Prop(t *testing.T) {
	ps := NewProp("s", "string", false)
	pi := NewProp("i", "int", false)
	pj := NewProp("i", "json.RawMessage", false)
	tests := []struct {
		name string
		prop Prop
		want Prop
	}{
		{"empty", nil, nil},
		{"string", ps, ps},
		{"string", pi, pi},
		{"string", pj, pj},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cn{
				p: tt.prop,
			}
			if got := c.Prop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cn_Type(t *testing.T) {
	tests := []struct {
		name string
		typ  Type
		want Type
	}{
		{"CondTypeAnd", CondTypeAnd, 1},
		{"CondTypeOr", CondTypeOr, 2},
		{"CondTypeEq", CondTypeEq, 3},
		{"CondTypeNeq", CondTypeNeq, 4},
		{"CondTypeLt", CondTypeLt, 5},
		{"CondTypeLte", CondTypeLte, 6},
		{"CondTypeGt", CondTypeGt, 7},
		{"CondTypeGte", CondTypeGte, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cn{
				t: tt.typ,
			}
			if got := c.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cn_Val(t *testing.T) {

	tests := []struct {
		name string
		val  interface{}
		want interface{}
	}{
		{"empty", nil, nil},
		{"int", 1, 1},
		{"string", "s", "s"},
		{"bool", true, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cn{
				v: tt.val,
			}
			if got := c.Val(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Val() = %v, want %v", got, tt.want)
			}
		})
	}
}
