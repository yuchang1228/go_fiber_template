package migrations

import (
	"app/config"
	"app/internal/models"
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateUsersTable, downCreateUsersTable)
}

func upCreateUsersTable(ctx context.Context, tx *sql.Tx) error {
	return config.DB_MIGRATOR.CreateTable(&models.User{})
}

func downCreateUsersTable(ctx context.Context, tx *sql.Tx) error {
	return config.DB_MIGRATOR.DropTable(&models.User{})
}
