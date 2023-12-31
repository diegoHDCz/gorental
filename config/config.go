package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	db, err = InitializePostgre()
	if err != nil {
		return fmt.Errorf("error intializing db: %v", err)
	}
	return nil
}

func GetPostgre() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
