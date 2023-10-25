package jwt

import (
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ramadoiranedar/rest-goswagger-api/internal/utils/constants"
)

// JWTMaker is an implementation of the Maker interface for creating and verifying JWT tokens
type JWTMaker struct {
	secretKey string
}

// CreateToken creates a new JWT token with user information and a specified duration
func (maker *JWTMaker) CreateToken(userID string, email, role string, duration time.Duration) (string, *PayloadJWT, error) {
	payload := NewPayload(userID, email, role, duration)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	return token, payload, err
}

// VerifyToken verifies the provided JWT token and returns the payload
func (maker *JWTMaker) VerifyToken(token string) (*PayloadJWT, error) {
	token = strings.TrimSpace(token)
	token = strings.ReplaceAll(token, "Bearer ", "") // remove the "Bearer " prefix if present

	// Define a key function for verifying the token's signature.
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, constants.JWT_ERR_INVALID_TOKEN // invalid token
		}
		return []byte(maker.secretKey), nil
	}

	// Parse the JWT token with the provided key function.
	jwtToken, err := jwt.ParseWithClaims(token, &PayloadJWT{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, constants.JWT_ERR_EXPIRED_TOKEN) {
			return nil, constants.JWT_ERR_EXPIRED_TOKEN // expired token
		}
		return nil, constants.JWT_ERR_FAILED_TOKEN // failed for other reasons
	}

	// Claim token JWT
	payload, ok := jwtToken.Claims.(*PayloadJWT)
	if !ok {
		return nil, constants.JWT_ERR_INVALID_PAYLOAD // invalid payload jwt
	}

	return payload, nil
}

// RefreshToken generates a new JWT token with an extended expiration time
// based on the existing token's payload
func RefreshToken(maker Maker, token string, duration time.Duration) (string, *PayloadJWT, error) {
	// Verify the existing token and extract its payload
	payload, err := maker.VerifyToken(token)
	if err != nil {
		return "", nil, err
	}
	payload.ExpiredAt = time.Now().Add(duration)

	// Create a new token with the existing payload and the extended duration
	newToken, newPayload, err := maker.CreateToken(payload.UserID, payload.Email, payload.Role, duration)
	if err != nil {
		return "", nil, err
	}

	return newToken, newPayload, nil
}

// RemoveToken invalidates a JWT token by creating a new token with a short duration,
// effectively logging the user out.
func RemoveToken(existingToken string, maker Maker) (string, error) {
	// Verify the existing token and extract its payload.
	payload, err := maker.VerifyToken(existingToken)
	if err != nil {
		return "", err
	}

	// Create a new token with the existing payload and a very short duration (e.g., 1 second).
	newToken, _, err := maker.CreateToken(payload.UserID, payload.Email, payload.Role, time.Second)
	if err != nil {
		return "", err
	}

	return newToken, nil
}
