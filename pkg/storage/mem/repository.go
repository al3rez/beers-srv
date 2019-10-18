package mem

import (
	"github.com/azbshiri/beers/pkg/adding"
)

// Memory storage keeps data in memory
type Storage struct {
	beers []adding.Beer
}

// Add saves the given beer to the repository
func (m *Storage) AddBeer(u adding.Beer) (*adding.Beer, error) {
	newB := adding.Beer{
		Id:   len(m.beers) + 1,
		Name: u.Name,
	}
	m.beers = append(m.beers, newB)

	return &newB, nil
}
