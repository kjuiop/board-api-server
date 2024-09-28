package member

import "board-api-server/internal/database"

type MemberRepository struct {
	db *database.MySqlClient
}

func NewUserMysqlRepository(db *database.MySqlClient) Repository {
	return &MemberRepository{db: db}
}

func (mr *MemberRepository) isExistByUsername(username string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (mr *MemberRepository) SignUp(data *MemberInfo) (int64, error) {
	//TODO implement me
	panic("implement me")
}
