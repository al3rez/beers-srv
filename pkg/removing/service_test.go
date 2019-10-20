package removing

import (
	"fmt"
	"reflect"
	"testing"
)

type dummyRepo struct {
	err   error
	beers []Beer
}

func (r *dummyRepo) RemoveBeer(b Beer) (*Beer, error) {
	if r.err != nil {
		return nil, r.err
	}

	for i, beer := range r.beers {
		if beer.Id == b.Id {
			r.beers[i] = r.beers[len(r.beers)-1]
			r.beers[len(r.beers)-1] = Beer{}
			r.beers = r.beers[:len(r.beers)-1]
			return &b, nil
		}
	}

	return &b, nil
}

func TestRemoveBeer(t *testing.T) {
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
		{"removes beer", NewService(&dummyRepo{err: nil, beers: []Beer{{0, "ali"}, {1, "boo"}}}), args{Beer{0, "ali"}}, &Beer{0, "ali"}, false},
		{"unexpected error", NewService(&dummyRepo{err: fmt.Errorf("pg cannot connect")}), args{Beer{0, "ali"}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.RemoveBeer(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.RemoveBeer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.RemoveBeer() = %v, want %v", got, tt.want)
			}
		})
	}
}
