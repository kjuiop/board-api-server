package member

type UserService struct {
}

func NewUserService() Service {
	return &UserService{}
}

func (u UserService) register(request *SignUpRequest) int64 {
	//TODO implement me
	panic("implement me")
}

func (u UserService) login(id, password string) UserInfo {
	//TODO implement me
	panic("implement me")
}

func (u UserService) isDuplicatedId(username string) bool {
	//TODO implement me
	panic("implement me")
}

func (u UserService) getUserInfo(username string) UserInfo {
	//TODO implement me
	panic("implement me")
}

func (u UserService) updateUserInfo(request *UpdateRequest) int64 {
	//TODO implement me
	panic("implement me")
}

func (u UserService) deleteByUsername(username string) int64 {
	//TODO implement me
	panic("implement me")
}
