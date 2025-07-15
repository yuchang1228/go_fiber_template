package database

import (
	"database/sql"
	"fmt"
	"strconv"

	"app/config"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GORM_DB *gorm.DB
var SQL_DB *sql.DB
var DB_MIGRATOR gorm.Migrator

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		panic("failed to parse database port")
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_HOST"),
		port,
		config.Config("DB_DATABASE"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{TranslateError: true})

	if err != nil {
		panic("failed to connect database")
	}

	GORM_DB = db
	SQL_DB, _ = db.DB()
	DB_MIGRATOR = db.Migrator()

	log.Info("Connected to database successfully")
}
