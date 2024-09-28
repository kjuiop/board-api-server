package member

import (
	"board-api-server/api/form"
	"fmt"
)

type UserService struct {
	repo Repository
}

func NewUserService(repo Repository) Service {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) SignUp(req SignUpRequest) (int64, error) {

	if err := us.CheckExistUsername(req.username); err != nil {
		return 0, err
	}

	return us.repo.SignUp(NewUserInfo(req))
}

func (us *UserService) CheckExistUsername(username string) error {
	isExist, err := us.repo.isExistByUsername(username)
	if err != nil {
		return fmt.Errorf("failed get username, err : %w", err)
	}
	if isExist {
		return form.GetCustomErr(form.ErrDuplicatedUsername)
	}

	return nil
}

func (us *UserService) login(id, password string) UserInfo {
	//TODO implement me
	panic("implement me")
}

func (us *UserService) isDuplicatedId(username string) bool {
	//TODO implement me
	panic("implement me")
}

func (us *UserService) getUserInfo(username string) UserInfo {
	//TODO implement me
	panic("implement me")
}

func (us *UserService) updateUserInfo(request *UpdateRequest) int64 {
	//TODO implement me
	panic("implement me")
}

func (us *UserService) deleteByUsername(username string) int64 {
	//TODO implement me
	panic("implement me")
}
