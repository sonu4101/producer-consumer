package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDatabase(dsn string) *sql.DB {
	// db, err := sql.Open("mysql", dsn)

	db, err := sql.Open(
		"mysql",
		"root:rootpass@tcp(mysql-service:3306)/producer_consumer",
	)
	fmt.Println("Sql error is: ", err)
	if err != nil {
		log.Fatal("failed to connect to MySQL:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("failed to ping MySQL:", err)
	}
	return db
}
