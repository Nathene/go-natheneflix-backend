package main

import (
	"backend/internal/repository"
	"backend/internal/repository/dbrepo"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const port = 3030

type app struct {
	DSN    string
	domain string
	DB     repository.DatabaseRepo
}

func main() {
	var app app
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	e := populateEnvVars()
	flag.StringVar(&app.DSN, "dsn", fmt.Sprintf(`
	host=%s port=%s user=%s password=%s dbname=%s sslmode=%s timezone=UTC connect_timeout=5`, e.host, e.port, e.user, e.password, e.dbname, e.sslmode),
		"Postgres Connection String")
	flag.Parse()

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	app.domain = "example.com"
	log.Printf("Listening on port %d...", port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}

}

type env struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
	sslmode  string
}

func populateEnvVars() *env {
	e := &env{}
	e.host = os.Getenv("PG_HOST")
	e.port = os.Getenv("PG_PORT")
	e.user = os.Getenv("PG_USER")
	e.password = os.Getenv("PG_PASS")
	e.dbname = os.Getenv("PG_NAME")
	e.sslmode = os.Getenv("PG_SSL")

	return e
}
