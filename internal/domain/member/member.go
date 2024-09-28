package member

import (
	"board-api-server/internal/domain/member/types"
	"time"
)

type SignUpRequest struct {
	username string
	password string
	nickname string
	isAdmin  bool
}

type SignUpRes struct {
	UserId int64
}

type UpdateRequest struct {
}

type MemberInfo struct {
	id         int64
	Username   string
	password   string
	nickname   string
	isAdmin    bool
	isWithDraw bool
	status     types.UserStatus
	createdAt  time.Time
	updatedAt  time.Time
}

func NewMemberInfo(req SignUpRequest) *MemberInfo {
	return &MemberInfo{
		Username:  req.username,
		password:  req.password,
		nickname:  req.nickname,
		isAdmin:   req.isAdmin,
		status:    types.DEFAULT,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}

type Service interface {
	SignUp(request SignUpRequest) (int64, error)
	login(id, password string) MemberInfo
	isDuplicatedId(username string) bool
	getUserInfo(username string) MemberInfo
	updateUserInfo(request *UpdateRequest) int64
	deleteByUsername(username string) int64
}

type Repository interface {
	isExistByUsername(username string) (bool, error)
	SignUp(data *MemberInfo) (int64, error)
}
