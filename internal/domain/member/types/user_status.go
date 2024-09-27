package types

type UserStatus int

const (
	DEFAULT UserStatus = iota
	ADMIN
	DELETED
)

func (s UserStatus) String() string {
	switch s {
	case DEFAULT:
		return "default"
	case ADMIN:
		return "admin"
	case DELETED:
		return "deleted"
	}

	return ""
}