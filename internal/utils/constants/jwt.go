package constants

import "errors"

const (
	JWT_MIN_SECRET_SIZE = 32
)

var (
	JWT_ERR_INVALID_TOKEN   = errors.New("invalid token")                        // the error for an invalid token
	JWT_ERR_FAILED_TOKEN    = errors.New("verification token for other reasons") // the error for an failed token
	JWT_ERR_EXPIRED_TOKEN   = errors.New("token has expired")                    // the error for an expired token
	JWT_ERR_INVALID_PAYLOAD = errors.New("invalid payload jwt")                  // the error for an invalid payload jwt
)
