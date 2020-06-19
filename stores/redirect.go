package stores

import "github.com/kousikmitra/shorty/models"

// RedirectStore is used to persist and find redirect urls
type RedirectStore interface {
	Find(code string) (*models.Redirect, error)
	Store(redirect *models.Redirect) error
}
