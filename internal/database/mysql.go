package database

import (
	"board-api-server/config"
	"database/sql"
	"fmt"
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
	return nil
}
