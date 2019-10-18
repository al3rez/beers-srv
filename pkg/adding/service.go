package adding

// Service provides beer authentication operations.
type Service interface {
	AddBeer(Beer) (*Beer, error)
}

// Repository provides access to beer repository.
type Repository interface {
	// AddBeer saves a given beer to the repository.
	AddBeer(Beer) (*Beer, error)
}

type service struct {
	Repo Repository
}

// NewService creates an authentication service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// AddBeer adds the given beer(s) to the database
func (s *service) AddBeer(b Beer) (*Beer, error) {
	// any validation can be done here
	beer, err := s.Repo.AddBeer(b)
	if err != nil {
		return nil, err
	}

	return beer, nil
}
