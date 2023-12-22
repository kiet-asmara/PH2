package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

func InitDb() *sql.DB {
	dbdsn := os.Getenv("DB_DSN")
	db, err := sql.Open("mysql", dbdsn)

	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
