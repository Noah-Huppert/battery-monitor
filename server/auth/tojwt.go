package auth

import (
	"fmt"

	"github.com/tj/robo/config"
)

// ToJWT converts an auth token to a JWT base64 encoded string. An application
// configuration must be provided to specify the finer details of the JWT
// creation process.
//
// An error will be returned if one occurs, nil on success.
func (t AuthToken) ToJWT(config *config.Config) (string, error) {
	// Get claims
	claims, err := t.Claims(config)
	if err != nil {
		return "", fmt.Errorf("error getting token claims: %s", err.Error())
	}

	// Make JWT
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if err != nil {
		return "", fmt.Errorf("error making JWT from token: %s", err.Error())
	}

	// Encode JWT to string
	str, err := token.SignedString(config.Auth.SigningSecret)
	if err != nil {
		return "", fmt.Errorf("error encoding JWT to string: %s", err.Error())
	}

	// Done
	return str, nil
}
