package member

import (
	"board-api-server/internal/domain/member/types"
	"time"
)

type SignUpRequest struct {
	username string
	password string
}

type UpdateRequest struct {
}

type UserInfo struct {
	id         int64
	username   string
	password   string
	nickname   string
	isAdmin    bool
	isWithDraw bool
	status     types.UserStatus
	createdAt  time.Time
	updatedAt  time.Time
}

type Service interface {
	register(request *SignUpRequest) int64
	login(id, password string) UserInfo
	isDuplicatedId(username string) bool
	getUserInfo(username string) UserInfo
	updateUserInfo(request *UpdateRequest) int64
	deleteByUsername(username string) int64
}

type Repository interface {
}
