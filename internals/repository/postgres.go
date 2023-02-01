package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	InitPSQLDBTables(db)

	return db, nil
}

func InitPSQLDBTables(db *sqlx.DB) error {
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS users(id SERIAL PRIMARY KEY, username TEXT, description TEXT, password TEXT)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	stmt, err = db.Prepare("CREATE TABLE IF NOT EXISTS posts(id SERIAL PRIMARY KEY, text TEXT, date timestamp, author_pk NUMERIC)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
