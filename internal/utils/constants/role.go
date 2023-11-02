package constants

const (
	ROLE_ID_ADMIN = uint64(1)
	ROLE_ID_USER  = uint64(2)
	ROLE_ID_GUEST = uint64(3)

	ROLE_SLUG_ADMIN = "admin"
	ROLE_SLUG_USER  = "user"
	ROLE_SLUG_GUEST = "guest"

	ROLE_NAME_ADMIN = "Admin"
	ROLE_NAME_USER  = "User"
	ROLE_NAME_GUEST = "Guest"
)

var (
	MAP_ROLE_ID = map[uint64]uint64{
		ROLE_ID_ADMIN: ROLE_ID_ADMIN,
		ROLE_ID_USER:  ROLE_ID_USER,
		ROLE_ID_GUEST: ROLE_ID_GUEST,
	}

	MAP_ROLE_SLUG = map[uint64]string{
		ROLE_ID_ADMIN: ROLE_SLUG_ADMIN,
		ROLE_ID_USER:  ROLE_SLUG_USER,
		ROLE_ID_GUEST: ROLE_SLUG_GUEST,
	}

	MAP_ROLE_NAME = map[uint64]string{
		ROLE_ID_ADMIN: ROLE_NAME_ADMIN,
		ROLE_ID_USER:  ROLE_NAME_USER,
		ROLE_ID_GUEST: ROLE_NAME_GUEST,
	}
)
