package member

import (
	"board-api-server/internal/domain/member/types"
	"time"
)

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	AdminYn  string `json:"admin_yn"`
}

type SignUpRes struct {
	UserId int64
}

type UpdateRequest struct {
}

type MemberInfo struct {
	id        int64
	Username  string
	password  string
	nickname  string
	adminYn   string
	status    types.UserStatus
	createdAt time.Time
	updatedAt time.Time
}

func NewMemberInfo(req SignUpRequest) *MemberInfo {
	return &MemberInfo{
		Username:  req.Username,
		password:  req.Password,
		nickname:  req.Nickname,
		adminYn:   req.AdminYn,
		status:    types.ACTIVE,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}

type Service interface {
	SignUp(request SignUpRequest) (int64, error)
	login(id, password string) MemberInfo
	isDuplicatedId(username string) error
	getUserInfo(username string) MemberInfo
	updateUserInfo(request *UpdateRequest) int64
	deleteByUsername(username string) int64
}

type Repository interface {
	isExistByUsername(username string) (bool, error)
	SignUp(data *MemberInfo) (int64, error)
}
