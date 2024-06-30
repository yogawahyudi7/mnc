package database

import (
	"fmt"
	"log"

	"github.com/yogawahyudi7/mnc/config"

	_ "github.com/lib/pq"
	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

func NewDatabase(config *config.Server) *gorm.DB {
	dns := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Password,
		config.Database.DBName,
		config.Database.SSLMode,
		config.Database.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	return db
}
