package main

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

func main() {
	_, err := migrate.New(
		"ngc-10/migrations",
		"postgres://postgres:12345@localhost:5432/ngc10?sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}

}
