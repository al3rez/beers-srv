package removing

// Service provides beer authentication operations.
type Service interface {
	RemoveBeer(Beer) (*Beer, error)
}

// Repository provides access to beer repository.
type Repository interface {
	// RemoveBeer saves a given beer to the repository.
	RemoveBeer(Beer) (*Beer, error)
}

type service struct {
	Repo Repository
}

// NewService creates an authentication service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// RemoveBeer adds the given beer(s) to the database
func (s *service) RemoveBeer(b Beer) (*Beer, error) {
	// any validation can be done here
	beer, err := s.Repo.RemoveBeer(b)
	if err != nil {
		return nil, err
	}

	return beer, nil
}
