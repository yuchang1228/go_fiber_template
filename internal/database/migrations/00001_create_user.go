package migrations

import (
	"app/config"
	"app/internal/model"
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateUser, downCreateUser)
}

func upCreateUser(ctx context.Context, tx *sql.Tx) error {
	return config.DB_MIGRATOR.CreateTable(&model.User{})
}

func downCreateUser(ctx context.Context, tx *sql.Tx) error {
	return config.DB_MIGRATOR.DropTable(&model.User{})
}
