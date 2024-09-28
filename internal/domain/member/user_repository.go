package member

import "board-api-server/internal/database"

type UserRepository struct {
	db *database.MySqlClient
}

func NewUserMysqlRepository(db *database.MySqlClient) Repository {
	return &UserRepository{db: db}
}
