package mem

import (
	"fmt"

	"github.com/azbshiri/beers-srv/pkg/adding"
	"github.com/azbshiri/beers-srv/pkg/removing"
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

// Add saves the given beer to the repository
func (m *Storage) RemoveBeer(b removing.Beer) (*removing.Beer, error) {
	for i, beer := range m.beers {
		if beer.Id == b.Id {
			m.beers[i] = m.beers[len(m.beers)-1]
			m.beers[len(m.beers)-1] = adding.Beer{}
			m.beers = m.beers[:len(m.beers)-1]
			return &b, nil
		}
	}

	return nil, fmt.Errorf("storage remove: cannot find\n")
}
