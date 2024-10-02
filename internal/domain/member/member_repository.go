package member

import (
	"board-api-server/internal/database"
	"strings"
)

type MemberRepository struct {
	db *database.MySqlClient
}

func NewUserMysqlRepository(db *database.MySqlClient) Repository {
	return &MemberRepository{db: db}
}

func (mr *MemberRepository) isExistByUsername(username string) (bool, error) {
	qs := query([]string{
		"SELECT",
		"COUNT(1)",
		"FROM board.member m",
		"WHERE 1=1",
		"AND m.username = ?",
		"AND m.status = 'active'",
	})

	count, err := mr.db.ExecCountQuery(qs, username)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (mr *MemberRepository) SignUp(data *MemberInfo) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func query(qs []string) string {
	return strings.Join(qs, " ") + ";"
}
