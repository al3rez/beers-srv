package adding

import (
	"errors"
	"testing"
)

type testRepo struct {
	returnError error
}

func (t *testRepo) AddBeer(u Beer) error {
	return t.returnError
}

func TestAddBeer(t *testing.T) {
	t.Run("test returns error", func(t *testing.T) {
		s := NewService(&testRepo{returnError: errors.New("bad")})
		u := Beer{}
		err := s.AddBeer(u)
		if err == nil {
			t.Fatal(err)
		}
	})

	t.Run("test adds beer", func(t *testing.T) {
		s := NewService(&testRepo{})
		u := Beer{}
		err := s.AddBeer(u)
		if err != nil {
			t.Fatal(err)
		}
	})
}
