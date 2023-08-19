package database

import (
	"database/sql"
	"fmt"
)

const (
	dbname   = "weather"
	user     = "postgres"
	password = "123456"
	host     = "localhost"
	port     = "5432"
	sslmode  = "disable"
)

func Connect() *sql.DB {
	// connect to database
	psglconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
	db, err := sql.Open("postgres", psglconn)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to Postgres")
	return db
}
