package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	dsn := "user:password@tcp(mariadb:3306)/cat_db"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	log.Println("Connected to DB")
	return db, nil
}
