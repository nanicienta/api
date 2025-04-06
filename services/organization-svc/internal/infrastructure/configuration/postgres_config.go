// Package configuration provides the configuration for the application
package configuration

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strconv"
	"time"
)

type dBConfig struct {
	databaseUser        string
	databasePassword    string
	databaseHost        string
	databasePort        int
	databaseName        string
	databaseMaxPoolSize int
	databaseMaxIdle     int
	databaseMaxLifetime int
	logging             bool
}

var dbConfiguration *dBConfig

func getDBConfig() *dBConfig {
	if dbConfiguration == nil {
		initDBConfig()
	}
	return dbConfiguration
}

func initDBConfig() {
	dbConfiguration = &dBConfig{}
	dbConfiguration.databaseUser = os.Getenv("DB_USER")
	dbConfiguration.databasePassword = os.Getenv("DB_PASS")
	dbConfiguration.databaseHost = os.Getenv("DB_HOST")
	var err error
	if dbConfiguration.databasePort, err = strconv.Atoi(os.Getenv("DB_PORT")); err != nil {
		panic(fmt.Sprintf("invalid DB_PORT: %v", err))
	}

	if dbConfiguration.databaseMaxLifetime, err = strconv.Atoi(os.Getenv("DB_POOL_MAX_LIFE_TIME")); err != nil {
		panic(fmt.Sprintf("invalid DB_POOL_MAX_LIFE_TIME: %v", err))
	}

	if dbConfiguration.databaseMaxPoolSize, err = strconv.Atoi(os.Getenv("DB_POOL_MAX")); err != nil {
		panic(fmt.Sprintf("invalid DB_POOL_MAX: %v", err))
	}

	if dbConfiguration.databaseMaxIdle, err = strconv.Atoi(os.Getenv("DB_POOL_MAX_IDLE")); err != nil {
		panic(fmt.Sprintf("invalid DB_POOL_MAX_IDLE: %v", err))
	}

	if dbConfiguration.logging, err = strconv.ParseBool(os.Getenv("DB_LOGGING")); err != nil {
		panic(fmt.Sprintf("invalid DB_LOGGING: %v", err))
	}
	dbConfiguration.databaseName = os.Getenv("DB_NAME")
}

// NewConnection Creates a new database connection
func NewConnection() (*gorm.DB, error) {
	dbConfig := getDBConfig()
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.databaseHost,
		dbConfig.databasePort,
		dbConfig.databaseUser,
		dbConfig.databasePassword,
		dbConfig.databaseName,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(dbConfig.databaseMaxPoolSize)
	sqlDB.SetMaxIdleConns(dbConfig.databaseMaxIdle)
	sqlDB.SetConnMaxLifetime(time.Duration(dbConfig.databaseMaxLifetime) * time.Minute)
	return db, nil
}
