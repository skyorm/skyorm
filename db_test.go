package skyorm

import (
	`context`
	`errors`
	`reflect`
	`testing`
)

func TestNewDB(t *testing.T) {
	type args struct {
		provider Provider
	}
	tests := []struct {
		name string
		args args
		want DB
	}{
		{
			name: "empty",
			args: args{},
			want: &db{},
		},
		{
			name: "demo",
			args: args{&demoProviderOK{}},
			want: &db{&demoProviderOK{}},
		},
		{
			name: "demoTest",
			args: args{&demoProviderOK{t}},
			want: &db{&demoProviderOK{t}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDB(tt.args.provider); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_db_Count(t *testing.T) {
	pOk := &demoProviderOK{t}
	pErr := &demoProviderErrored{t}
	type fields struct {
		p Provider
	}
	type args struct {
		store     Store
		condition Cond
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "ok",
			fields:  fields{pOk},
			args:    args{demoStore, nil},
			want:    1,
			wantErr: false,
		},
		{
			name:    "ok",
			fields:  fields{pErr},
			args:    args{demoStore, nil},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &db{
				p: tt.fields.p,
			}
			got, err := d.Count(tt.args.store, tt.args.condition)
			if (err != nil) != tt.wantErr {
				t.Errorf("Count() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Count() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_db_CountContext(t *testing.T) {
	pOk := &demoProviderOK{t}
	pErr := &demoProviderErrored{t}
	type fields struct {
		p Provider
	}
	type args struct {
		ctx       context.Context
		store     Store
		condition Cond
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "ok",
			fields:  fields{pOk},
			args:    args{context.Background(), demoStore, nil},
			want:    1,
			wantErr: false,
		},
		{
			name:    "ok",
			fields:  fields{pErr},
			args:    args{context.Background(), demoStore, nil},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &db{
				p: tt.fields.p,
			}
			got, err := d.CountContext(tt.args.ctx, tt.args.store, tt.args.condition)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CountContext() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_db_Delete(t *testing.T) {
	pOk := &demoProviderOK{t}
	pErr := &demoProviderErrored{t}
	type fields struct {
		p Provider
	}
	type args struct {
		store     Store
		condition Cond
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "ok",
			fields:  fields{pOk},
			args:    args{demoStore, nil},
			wantErr: false,
		},
		{
			name:    "ok",
			fields:  fields{pErr},
			args:    args{demoStore, nil},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &db{
				p: tt.fields.p,
			}
			if err := d.Delete(tt.args.store, tt.args.condition); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_db_DeleteContext(t *testing.T) {
	pOk := &demoProviderOK{t}
	pErr := &demoProviderErrored{t}
	type fields struct {
		p Provider
	}
	type args struct {
		ctx       context.Context
		store     Store
		condition Cond
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "ok",
			fields:  fields{pOk},
			args:    args{context.Background(), demoStore, nil},
			wantErr: false,
		},
		{
			name:    "ok",
			fields:  fields{pErr},
			args:    args{context.Background(), demoStore, nil},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &db{
				p: tt.fields.p,
			}
			if err := d.DeleteContext(tt.args.ctx, tt.args.store, tt.args.condition); (err != nil) != tt.wantErr {
				t.Errorf("DeleteContext() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_db_ErrNotFound(t *testing.T) {
	pOk := &demoProviderOK{t}
	pErr := &demoProviderErrored{t}
	type fields struct {
		p Provider
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "ok",
			fields:  fields{pOk},
			wantErr: true,
		},
		{
			name:    "ok",
			fields:  fields{pErr},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &db{
				p: tt.fields.p,
			}
			if err := d.ErrNotFound(); (err != nil) != tt.wantErr {
				t.Errorf("ErrNotFound() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_db_Find(t *testing.T) {
	pOk := &demoProviderOK{t}
	pErr := &demoProviderErrored{t}
	type fields struct {
		p Provider
	}
	type args struct {
		store     Store
		condition Cond
		limit     int
		offset    int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Model
		wantErr bool
	}{
		{
			name:    "ok",
			fields:  fields{pOk},
			args:    args{demoStore, nil, 1, 1},
			want:    []Model{},
			wantErr: false,
		},
		{
			name:    "err",
			fields:  fields{pErr},
			args:    args{demoStore, nil, 1, 1},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &db{
				p: tt.fields.p,
			}
			got, err := d.Find(tt.args.store, tt.args.condition, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_db_FindContext(t *testing.T) {
	pOk := &demoProviderOK{t}
	pErr := &demoProviderErrored{t}
	type fields struct {
		p Provider
	}
	type args struct {
		ctx       context.Context
		store     Store
		condition Cond
		limit     int
		offset    int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Model
		wantErr bool
	}{
		{
			name:    "ok",
			fields:  fields{pOk},
			args:    args{context.Background(), demoStore, nil, 1, 1},
			want:    []Model{},
			wantErr: false,
		},
		{
			name:    "err",
			fields:  fields{pErr},
			args:    args{context.Background(), demoStore, nil, 1, 1},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &db{
				p: tt.fields.p,
			}
			got, err := d.FindContext(tt.args.ctx, tt.args.store, tt.args.condition, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindContext() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_db_IsNotFound(t *testing.T) {
	pOk := &demoProviderOK{t}
	pErr := &demoProviderErrored{t}
	type fields struct {
		p Provider
	}
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "nil",
			fields: fields{pOk},
			args:   args{nil},
			want:   false,
		},
		{
			name:   "ok",
			fields: fields{pOk},
			args:   args{demoProviderErr},
			want:   true,
		},
		{
			name:   "err",
			fields: fields{pErr},
			args:   args{errors.New("err")},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &db{
				p: tt.fields.p,
			}
			if got := d.IsNotFound(tt.args.err); got != tt.want {
				t.Errorf("IsNotFound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_db_Populate(t *testing.T) {
	pOk := &demoProviderOK{t}
	pErr := &demoProviderErrored{t}
	type fields struct {
		p Provider
	}
	type args struct {
		model Model
		pk    interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "ok",
			fields:  fields{pOk},
			args:    args{demoStore.Model(), 1},
			wantErr: false,
		},
		{
			name:    "err",
			fields:  fields{pErr},
			args:    args{demoStore.Model(), 1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &db{
				p: tt.fields.p,
			}
			if err := d.Populate(tt.args.model, tt.args.pk); (err != nil) != tt.wantErr {
				t.Errorf("Populate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_db_PopulateContext(t *testing.T) {
	pOk := &demoProviderOK{t}
	pErr := &demoProviderErrored{t}
	type fields struct {
		p Provider
	}
	type args struct {
		ctx   context.Context
		model Model
		pk    interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "ok",
			fields:  fields{pOk},
			args:    args{context.Background(), demoStore.Model(), 1},
			wantErr: false,
		},
		{
			name:    "err",
			fields:  fields{pErr},
			args:    args{context.Background(), demoStore.Model(), 1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &db{
				p: tt.fields.p,
			}
			if err := d.PopulateContext(tt.args.ctx, tt.args.model, tt.args.pk); (err != nil) != tt.wantErr {
				t.Errorf("PopulateContext() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_db_Put(t *testing.T) {
	pOk := &demoProviderOK{t}
	pErr := &demoProviderErrored{t}
	type fields struct {
		p Provider
	}
	type args struct {
		models []Model
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "ok1",
			fields:  fields{pOk},
			args:    args{[]Model{demoStore.Model()}},
			wantErr: false,
		},
		{
			name:    "ok2",
			fields:  fields{pOk},
			args:    args{[]Model{demoStore.Model(), demoStore.Model()}},
			wantErr: false,
		},
		{
			name:    "err1",
			fields:  fields{pErr},
			args:    args{[]Model{demoStore.Model()}},
			wantErr: true,
		},
		{
			name:    "err1",
			fields:  fields{pErr},
			args:    args{[]Model{demoStore.Model(), demoStore.Model()}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &db{
				p: tt.fields.p,
			}
			if err := d.Put(tt.args.models...); (err != nil) != tt.wantErr {
				t.Errorf("Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_db_PutContext(t *testing.T) {
	pOk := &demoProviderOK{t}
	pErr := &demoProviderErrored{t}
	type fields struct {
		p Provider
	}
	type args struct {
		ctx    context.Context
		models []Model
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "no models",
			fields:  fields{pOk},
			args:    args{ctx: context.Background()},
			wantErr: true,
		},
		{
			name:    "ok1",
			fields:  fields{pOk},
			args:    args{context.Background(), []Model{demoStore.Model()}},
			wantErr: false,
		},
		{
			name:    "ok2",
			fields:  fields{pOk},
			args:    args{context.Background(), []Model{demoStore.Model(), demoStore.Model()}},
			wantErr: false,
		},
		{
			name:    "err1",
			fields:  fields{pErr},
			args:    args{context.Background(), []Model{demoStore.Model()}},
			wantErr: true,
		},
		{
			name:    "err1",
			fields:  fields{pErr},
			args:    args{context.Background(), []Model{demoStore.Model(), demoStore.Model()}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &db{
				p: tt.fields.p,
			}
			if err := d.PutContext(tt.args.ctx, tt.args.models...); (err != nil) != tt.wantErr {
				t.Errorf("PutContext() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_db_Update(t *testing.T) {
	p := NewProp("s", "string", false)
	pOk := &demoProviderOK{t}
	pErr := &demoProviderErrored{t}
	type fields struct {
		p Provider
	}
	type args struct {
		store     Store
		condition Cond
		values    []Val
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "ok",
			fields:  fields{pOk},
			args:    args{demoStore, nil, []Val{NewVal(p, 1)}},
			wantErr: false,
		},
		{
			name:    "err",
			fields:  fields{pErr},
			args:    args{demoStore, nil, []Val{NewVal(p, 1)}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &db{
				p: tt.fields.p,
			}
			if err := d.Update(tt.args.store, tt.args.condition, tt.args.values...); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_db_UpdateContext(t *testing.T) {
	p := NewProp("s", "string", false)
	pOk := &demoProviderOK{t}
	pErr := &demoProviderErrored{t}
	type fields struct {
		p Provider
	}
	type args struct {
		ctx       context.Context
		store     Store
		condition Cond
		values    []Val
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "ok",
			fields:  fields{pOk},
			args:    args{context.Background(), demoStore, nil, []Val{NewVal(p, 1)}},
			wantErr: false,
		},
		{
			name:    "err",
			fields:  fields{pErr},
			args:    args{context.Background(), demoStore, nil, []Val{NewVal(p, 1)}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &db{
				p: tt.fields.p,
			}
			if err := d.UpdateContext(tt.args.ctx, tt.args.store, tt.args.condition, tt.args.values...); (err != nil) != tt.wantErr {
				t.Errorf("UpdateContext() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
