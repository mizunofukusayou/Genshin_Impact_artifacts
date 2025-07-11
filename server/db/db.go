package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func New() (*DB, error) {
	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		return nil, fmt.Errorf("DB_NAME environment variable not set")
	}

	host, ok := os.LookupEnv("DB_HOSTNAME")
	if !ok {
		return nil, fmt.Errorf("DB_HOSTNAME environment variable not set")
	}

	port, ok := os.LookupEnv("DB_PORT")
	if !ok {
		return nil, fmt.Errorf("DB_PORT environment variable not set")
	}

	user, ok := os.LookupEnv("DB_USER")
	if !ok {
		return nil, fmt.Errorf("DB_USER environment variable not set")
	}

	password, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		return nil, fmt.Errorf("DB_PASSWORD environment variable not set")
	}

	// GORMのDSN形式で接続文字列を作成
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &DB{db: db}, nil
}
