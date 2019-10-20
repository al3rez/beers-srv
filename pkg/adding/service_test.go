package adding

import (
	"reflect"
	"testing"
)

type dummyRepo struct {
	err error
}

func (r *dummyRepo) AddBeer(b Beer) (*Beer, error) {
	if r.err != nil {
		return nil, r.err
	}

	return &b, nil
}

func TestAddBeer(t *testing.T) {
	type args struct {
		b Beer
	}
	tests := []struct {
		name    string
		s       Service
		args    args
		want    *Beer
		wantErr bool
	}{
		{"adds beer", NewService(&dummyRepo{nil}), args{Beer{}}, &Beer{}, false},
		{"unexpected error", NewService(&dummyRepo{nil}), args{}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.AddBeer(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.AddBeer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.AddBeer() = %v, want %v", got, tt.want)
			}
		})
	}
}
