package util

import (
    "fmt"
    "log"
    "sync"

    "github.com/pp00x/hydrate/config"
    "github.com/pp00x/hydrate/internal/model"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var (
    db   *gorm.DB
    once sync.Once
)

func GetDB() *gorm.DB {
    once.Do(func() {
        dbConfig := config.AppConfig.Database
        dsn := fmt.Sprintf(
            "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
            dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.Port, dbConfig.SSLMode, dbConfig.TimeZone,
        )
        var err error
        db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err != nil {
            log.Fatalf("Failed to connect to database: %v", err)
        }

        // Auto-migrate models
        err = db.AutoMigrate(&model.User{}, &model.WaterIntake{})
        if err != nil {
            log.Fatalf("Failed to migrate database: %v", err)
        }
    })
    return db
}