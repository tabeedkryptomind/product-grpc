package database

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "admin"
	password = "admin123"
	dbname   = "postgresdb"
)

func PostgressConnect() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatalf("error on connecting database postgress %v", err)
		return nil, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatalf("error on Ping postgress %v", err)
		return nil, err
	}
	return db, nil
}
