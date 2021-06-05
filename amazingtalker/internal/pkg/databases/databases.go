package databases

import (
	"time"

	"github.com/blackhorseya/amazingtalker/internal/pkg/config"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"

	// import mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// NewMariaDB init mariadb client
func NewMariaDB(cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", cfg.DB.URL)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewMariaDB)
