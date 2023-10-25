package jwt

import (
	"fmt"
	"time"

	"github.com/ramadoiranedar/rest-goswagger-api/internal/utils/constants"
)

// Maker is an interface for creating and verifying JWT tokens
type Maker interface {
	CreateToken(userID string, email, role string, duration time.Duration) (string, *PayloadJWT, error)
	VerifyToken(token string) (*PayloadJWT, error)
}

// NewJWTMaker creates a new JWTMaker instance with a secret key
func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < constants.JWT_MIN_SECRET_SIZE {
		return nil, fmt.Errorf("invalid secret key size, must be at least %d characters", constants.JWT_MIN_SECRET_SIZE)
	}
	return &JWTMaker{secretKey}, nil
}
