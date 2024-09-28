package member

import "board-api-server/internal/database"

type UserRepository struct {
	db *database.MySqlClient
}

func NewUserMysqlRepository(db *database.MySqlClient) Repository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) isExistByUsername(username string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (ur *UserRepository) SignUp(data *UserInfo) (int64, error) {
	//TODO implement me
	panic("implement me")
}
