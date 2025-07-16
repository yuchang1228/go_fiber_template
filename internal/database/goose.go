package database

import (
	"app/config"

	"github.com/pressly/goose/v3"
)

func Migrate() {
	if err := goose.SetDialect(config.Config("DB_DRIVER")); err != nil {
		panic(err)
	}

	if err := goose.Up(config.SQL_DB, "internal/database/migrations"); err != nil {
		panic(err)
	}
}
