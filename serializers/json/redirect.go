package json

import (
	"encoding/json"

	"github.com/kousikmitra/shorty/models"
	"github.com/pkg/errors"
)

// Redirect json serializer type
type Redirect struct{}

// Encode a redirect object
func (r *Redirect) Encode(input *models.Redirect) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializers.Redirect.Encode")
	}
	return rawMsg, nil
}

// Decode a redirect byte
func (r *Redirect) Decode(input []byte) (*models.Redirect, error) {
	redirect := &models.Redirect{}
	if err := json.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serializers.Redirect.Decode")
	}
	return redirect, nil
}
