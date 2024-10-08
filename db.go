package main

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDB(s string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", s)
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic("Failed to ping the database: " + err.Error())
	}

	db.DB.SetMaxOpenConns(200)
	db.DB.SetMaxIdleConns(30)
	db.DB.SetConnMaxLifetime(10 * time.Minute)

	return db
}
