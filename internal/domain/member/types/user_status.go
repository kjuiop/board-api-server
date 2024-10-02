package types

type UserStatus int

const (
	ACTIVE UserStatus = iota
	INACTIVE
	WITHDRAW
)

func (s UserStatus) String() string {
	switch s {
	case ACTIVE:
		return "active"
	case INACTIVE:
		return "inactive"
	case WITHDRAW:
		return "withdraw"
	}

	return ""
}
