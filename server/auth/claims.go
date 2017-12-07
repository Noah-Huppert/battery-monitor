package auth

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

// Claims returns the JWT claims the Auth Token model provides
func (t AuthToken) Claims() map[string]interface{} {

}
