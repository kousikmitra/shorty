package services

import (
	"time"

	"github.com/kousikmitra/shorty/models"
	"github.com/kousikmitra/shorty/stores"
	"github.com/pkg/errors"
	"github.com/teris-io/shortid"
	"gopkg.in/dealancer/validate.v2"
)

// RedirectService is used to store and find redirect urls
type RedirectService interface {
	Find(code string) (*models.Redirect, error)
	Store(redirect *models.Redirect) error
}

type redirectService struct {
	redirectStore stores.RedirectStore
}

// NewRedirectService will return a new redirect service
func NewRedirectService(redirectStore stores.RedirectStore) RedirectService {
	return &redirectService{
		redirectStore,
	}
}

func (s *redirectService) Find(code string) (*models.Redirect, error) {
	return s.redirectStore.Find(code)
}

func (s *redirectService) Store(redirect *models.Redirect) error {
	if err := validate.Validate(redirect); err != nil {
		return errors.Wrap(models.ErrRedirectInvalid, "service.Redirect.Store")
	}

	redirect.Code = shortid.MustGenerate()
	redirect.CreatedAt = time.Now().UTC().Unix()
	return s.redirectStore.Store(redirect)
}
