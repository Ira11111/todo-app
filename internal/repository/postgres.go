package repository

import (
	"fmt"
	"github.com/Ira11111/todo-app/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB(cfg *config.Config) (*sqlx.DB, error) {
	dbCfg := &cfg.Database
	db, err := sqlx.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbCfg.Host, dbCfg.Port, dbCfg.Username, dbCfg.DBName, dbCfg.Password, dbCfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
