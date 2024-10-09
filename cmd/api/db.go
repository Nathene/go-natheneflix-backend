package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func openDB(dsn string) (*sql.DB, error) {
	return sql.Open("pgx", dsn)
}

func (a *app) connectToDB() (*sql.DB, error) {
	conn, err := openDB(a.DSN)
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to Postgres")
	return conn, nil
}
