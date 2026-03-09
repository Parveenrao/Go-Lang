package database

import (
	"go_crud_api/internals/config"
	"go_crud_api/internals/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


func Connect(cfg *config.Config) (*gorm.DB , error) {

	logLevel := logger.Silent

	if cfg.AppEnv == "development" {
		logLevel = logger.Info
	}

	db , err := gorm.Open(postgres.Open(cfg.DB.DSN()) , &gorm.Config{

		Logger: logger.Default.LogMode(logLevel),
	})

	if err != nil {
		return  nil , err
	}


	// connection pool setting

	sqlDB  , err := db.DB()

	if err != nil {
		return  nil , err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)


	log.Println("Database connected successfully")
	return  db ,nil



}



func Automigrate(db *gorm.DB) error {
	log.Println("Running database migration")

	return  db.AutoMigrate(&models.Product{},
	)

}
