package infrastructure

import (
	"fmt"
	"time"

	"github.com/didikz/goshu/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDB(dbConfig config.Database) *sqlx.DB {
	db, err := sqlx.Connect(dbConfig.Driver, fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=%s", dbConfig.Driver, dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBname, dbConfig.SSLMode))
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	db.DB.SetMaxOpenConns(200)
	db.DB.SetMaxIdleConns(30)
	db.DB.SetConnMaxLifetime(10 * time.Minute)

	return db
}
