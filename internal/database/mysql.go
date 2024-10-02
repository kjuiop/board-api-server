package database

import (
	"board-api-server/config"
	"board-api-server/internal/utils"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

type MySqlClient struct {
	cfg config.Mysql
	db  *sql.DB
}

func NewMysqlClient(cfg config.Mysql) (*MySqlClient, error) {

	url := fmt.Sprintf("%s:%s@tcp(%s)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Database)
	db, err := sql.Open(cfg.Driver, url)
	if err != nil {
		return nil, fmt.Errorf("failed connect mysql, err : %w", err)
	}

	client := &MySqlClient{
		cfg: cfg,
		db:  db,
	}

	if err := client.checkDefaultTable(); err != nil {
		return nil, err
	}

	return client, nil
}

func (m *MySqlClient) checkDefaultTable() error {

	query := checkExistChatQuery()

	var count int
	if err := m.db.QueryRow(query).Scan(&count); err != nil {
		log.Fatal("Error checking table existence: ", err)
	}

	if count > 0 {
		return nil
	}

	content, err := utils.ReadFileContent("/internal/database/schema/database.sql")
	if err != nil {
		log.Fatalf("error opening SQL file: %v", err)
	}

	queries := strings.Split(content, ";")

	for _, query := range queries {
		query = strings.TrimSpace(query)
		if query == "" {
			continue
		}

		_, err = m.db.Exec(query)
		if err != nil {
			return fmt.Errorf("error executing query : %s, err : %w", query, err)
		}
	}

	return nil
}

func (m *MySqlClient) ExecCountQuery(query string, args ...interface{}) (int, error) {
	var count int

	// 쿼리 실행 및 스캔
	if err := m.db.QueryRow(query, args...).Scan(&count); err != nil {
		return 0, fmt.Errorf("failed to execute count query: %w", err)
	}

	return count, nil
}

func checkExistChatQuery() string {
	return `
    SELECT COUNT(*)
    FROM information_schema.tables
    WHERE table_schema = 'board' AND table_name = 'member';
    `
}
