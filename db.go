package main

import (
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

	return db
}
