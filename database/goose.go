package database

import (
	"app/config"

	"github.com/pressly/goose/v3"
)

// asjkdias
func Migrate() {
	if err := goose.SetDialect(config.Config("DB_DRIVER")); err != nil {
		panic(err)
	}

	if err := goose.Up(SQL_DB, "database/migrations"); err != nil {
		panic(err)
	}
}
