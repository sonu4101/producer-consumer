package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDatabase(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("failed to connect to MySQL:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("failed to ping MySQL:", err)
	}
	return db
}
