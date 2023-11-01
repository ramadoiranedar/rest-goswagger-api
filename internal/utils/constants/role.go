package constants

const (
	SLUG_ROLE_ADMIN = "admin"
	SLUG_ROLE_USER  = "user"

	ROLE_ID_ADMIN = uint64(1)
	ROLE_ID_USER  = uint64(2)
)

var (
	MAP_ROLE_USER = map[string]uint64{
		SLUG_ROLE_ADMIN: ROLE_ID_ADMIN,
		SLUG_ROLE_USER:  ROLE_ID_USER,
	}
)
