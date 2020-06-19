package serializer

import "github.com/kousikmitra/shorty/models"

// RedirectSerializer serializes Redirect model
type RedirectSerializer interface {
	Encode(redirect *models.Redirect) ([]byte, error)
	Decode(input []byte) (*models.Redirect, error)
}
