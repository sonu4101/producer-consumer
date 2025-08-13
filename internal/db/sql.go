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
	 _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS messages (
            id BIGINT AUTO_INCREMENT PRIMARY KEY,
            message VARCHAR(255) NOT NULL,
            created_at DATETIME NOT NULL
        )
    `)
    if err != nil {
        log.Fatal("failed to create messages table:", err)
    }
	return db
}
