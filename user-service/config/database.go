package config

import (
	"fmt"
	"net/url"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	encodedPassword := url.QueryEscape(Config.Database.Password)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		Config.Database.Host, Config.Database.Username, encodedPassword, Config.Database.Name, Config.Database.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(Config.Database.MaxIdleConnection)
	sqlDB.SetMaxOpenConns(Config.Database.MaxOpenConnection)
	sqlDB.SetConnMaxLifetime(time.Duration(Config.Database.MaxLifetimeConnection) * time.Second)
	sqlDB.SetConnMaxIdleTime(time.Duration(Config.Database.MaxIdleTime))

	return db, nil
}
