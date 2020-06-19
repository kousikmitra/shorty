package memory

import (
	"github.com/kousikmitra/shorty/models"
	"github.com/kousikmitra/shorty/stores"
	"github.com/pkg/errors"
)

type memoryStore struct {
	maxSize int
	storage []models.Redirect
}

// NewMemoryStore return a new memory store
func NewMemoryStore(maxSize int) (stores.RedirectStore, error) {
	store := &memoryStore{
		maxSize: maxSize,
		storage: make([]models.Redirect, maxSize),
	}

	return store, nil
}

func (s *memoryStore) Find(code string) (*models.Redirect, error) {
	for _, redirect := range s.storage {
		if redirect.Code == code {
			return &redirect, nil
		}
	}
	return nil, errors.Wrap(models.ErrRedirectNotFound, "stores.Redirect.Find")
}

func (s *memoryStore) Store(redirect *models.Redirect) error {
	s.storage = append(s.storage, *redirect)
	return nil
}
