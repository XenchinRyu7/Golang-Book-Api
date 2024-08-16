package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/db_book"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Check if the connection is successful
	if err = db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database successfully!")
	return db, nil
}
