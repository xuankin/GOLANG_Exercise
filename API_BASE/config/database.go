package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(config *Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.DbSource), &gorm.Config{})
	if err != nil {
		log.Fatal("Khong the ket noi den database :", err)
	}
	log.Println("Ket noi den database thanh cong")
	return db
}
