package config

import (
	"github.com/diegoHDCz/gorental/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgre() (*gorm.DB, error) {
	logger := GetLogger("postgre")

	dsn := "host=localhost user=postgre password=postgre dbname=gorental port=5432 sslmode=disable"
	logger.Debugf("connect: %v", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("postgre op error %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Category{})
	if err != nil {
		logger.Errorf("postgre migrate error %v", err)
		return nil, err
	}
	return db, nil
}
