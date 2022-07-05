package config

import (
	"blog-api/src/model"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getDsn() string {
	var (
		username = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASS")
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		database = os.Getenv("DB_NAME")
	)

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, database,
	)

	return dsn
}

func CreateConnection() (*gorm.DB, error) {
	gormConfig := &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent),
	}

	db, err := gorm.Open(mysql.Open(getDsn()), gormConfig)

	db.AutoMigrate(
		&model.User{},
		&model.Post{},
		&model.Category{},
	)

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()

	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(10000)
	sqlDB.SetMaxIdleConns(10000)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)

	return db, nil
}
