package jwt

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/ramadoiranedar/rest-goswagger-api/internal/utils/constants"
)

// PayloadJWT represents the data stored in a JWT token
type PayloadJWT struct {
	UserID    string       `json:"user_id"`
	Email     strfmt.Email `json:"email"`
	Role      string       `json:"role"`
	IssuedAt  time.Time    `json:"issued_at"`
	ExpiredAt time.Time    `json:"expired_at"`
}

// NewPayload creates a new Payload with user information and duration
func NewPayload(income *PayloadJWT, duration time.Duration) *PayloadJWT {
	return &PayloadJWT{
		UserID:    income.UserID,
		Email:     income.Email,
		Role:      income.Role,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
}

// Valid checks if the token is still valid (not expired)
// issues: cannot use payload (variable of type *PayloadJWT) as "github.com/dgrijalva/jwt-go".Claims value in argument to jwt.NewWithClaims: *PayloadJWT does not implement "github.com/dgrijalva/jwt-go".Claims (missing method Valid)
func (payload *PayloadJWT) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return constants.JWT_ERR_EXPIRED_TOKEN
	}
	return nil
}
