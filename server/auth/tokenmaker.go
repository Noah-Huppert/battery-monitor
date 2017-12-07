package auth

import "github.com/dgrijalva/jwt-go"

// TokenMaker is an interface which wraps the method used to generate auth
// tokens
type TokenMaker interface {
	// MakeToken generates a JWT with the specified arguments. It returns
	// an AuthToken model, recording metadata about the token, and the
	// JWT base64 encoded text
	//
	// Args:
	// - userID (uint): ID of user to issue auth token to
	// - deviceID (uint): ID of the device to issue an auth token to
	//
	// Returns:
	// - AuthToken: Model holding metadata about auth token
	// - string: JWT encoded in string form
	// - error: If an error occurred, nil on success
	MakeToken(userID uint, deviceID uint) (AuthToken, string, error)
}

// DefaultTokenMaker is the implementation of the TokenMaker interface used
// commonly throughout the program
type DefaultTokenMaker struct {
	// db used to query persistence layer
	db *db.Db

	// config holds application configuration values
	config Config
}

// NewTokenMaker creates a new Token Maker with the default
// implementation.
//
// Args:
// - db (*db.Db): Database used to query models
//
// Returns:
// - TokenMaker: Default implementation of TokenMaker
func NewTokenMaker(config Config) TokenMaker {
	// Make db connection
	return DefaultTokenMaker{db}
}

// MakeToken creates a new JWT and AuthToken model
func (m DefaultTokenMaker) MakeToken(userID uint, deviceID uint) (AuthToken, string, error) {
	// Default auth token and jwt values to return on err
	var AuthToken token
	var jwt string

	// Check device with id, belonging to specified user, exists
	var Device *device
	m.db.Db.Where(&Device{ID: deviceID, UserID: userID}).First(device)
	if err := m.db.Db.Error; err != nil {
		return token, jwt, fmt.Errorf("error checking user and device ID: %s",
			err.Error())
	}

	if device == nil {
		return token, jwt, fmt.Errorf("no device with specified id, " +
			"belonging to specifeid user exists")
	}

	// If device exists and belongs to user
	// TODO make jwt
}
