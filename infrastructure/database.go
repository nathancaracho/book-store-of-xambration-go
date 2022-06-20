package infrastructure

import (
	"book-store-of-xambration-go/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection() (db *gorm.DB, err error) {
	dsn := "host=localhost user=postgres password=postgres dbname=book_store port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return
}

func Migrate() error {
	db, err := GetConnection()
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&models.Book{})
	return err
}
