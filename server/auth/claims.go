package auth

import (
	"fmt"
	"regexp"
)

// IssuerClaimKey holds the issuer field's key in a JWT claim
const IssuerClaimKey string = "iss"

// SubjectClaimKey holds the subject field's key in a JWT claim
const SubjectClaimKey string = "sub"

// AudienceClaimKey holds the audience field's key in a JWT claim
const AudienceClaimKey string = "aud"

// IssuedAtClaimKey holds the issued at field's key in a JWT claim
const IssuedAtClaimKey string = "iat"

// IDClaimKey holds the JWT ID field's key in a claim
const IDClaimKey string = "jti"

const SubjectClaimRegex regexp.Regexp = regexp.MustCompile("^([0-9]+):([0-9]+)$")

// makeSub combines the UserID and DeviceID fields into one single string,
// which will be used in the subject JWT claim. In the form:
// 	<user id>:<device id>
// An error is returned if the AuthToken is not valid
func (t AuthToken) makeSub() (string, error) {
	// Check valid
	if err := t.Validate(); err != nil {
		return "", fmt.Errorf("error validating auth token: %s", err.Error())
	}

	// Format
	return fmt.Sprintf("%s:%s", t.UserID, t.DeviceID), nil
}

// parseSub reads a JWT subject claim and saves the user and device ID into the
// token struct. Returns an error if one occurs parsing the subject claim.
func (t AuthToken) parseSub(sub string) error {
	// Use regex to extract
	ids := SubjectClaimRegex.FindStringSubmatch(sub)
	if ids == nil {
		return fmt.Errorf("error extracting ids, incorrect format")
	}

	// Get both ids
	ids = ids[1:]

	t.UserID, err = strconv.ParseUInt(ids[0], 10, 64)
	if err != nil {
		return fmt.Errorf("error parsing user id into uint: %s", err.Error())
	}

	t.DeviceID, err = strconv.ParseUInt(ids[1], 10, 64)
	if err != nil {
		return fmt.Errorf("error parsing device id into uint: %s", err.Error())
	}

	// All good
	return nil
}

// Claims returns the JWT claims the Auth Token model provides
//
// Args:
// 	- config: App configuration
func (t AuthToken) Claims(config *config.Config) (map[string]interface{}, error) {
	// Validate token
	if err := t.Validate(); err != nil {
		return fmt.Errorf("error validating token: %s", err.Error())
	}

	// Assemble claims
	c = make(map[string]interface{})

	// ID of token
	c[IDClaimKey] = t.ID

	// Issuer
	c[IssuerClaimKey] = config.Auth.Identity

	// Subject
	c[SubjectClaimKey], err = t.makeSub()
	if err != nil {
		return c, fmt.Errorf("error assembling token subject: %s", err.Error())
	}

	// Audience
	c[AudienceClaimKey] = config.Auth.Identity

	// Issued at
	c[IssuedAtClaimKey] = t.CreatedAt

	// Success
	return c, nil
}
