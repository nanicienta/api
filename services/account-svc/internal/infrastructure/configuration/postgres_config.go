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
	var err error
	dbConfiguration = &dBConfig{}
	dbConfiguration.databaseUser = os.Getenv("DB_USER")
	dbConfiguration.databasePassword = os.Getenv("DB_PASS")
	dbConfiguration.databaseHost = os.Getenv("DB_HOST")
	dbConfiguration.databasePort, err = strconv.Atoi(os.Getenv("DB_PORT"))
	dbConfiguration.databaseMaxLifetime, err = strconv.Atoi(os.Getenv("DB_POOL_MAX_LIFE_TIME"))
	dbConfiguration.databaseMaxPoolSize, err = strconv.Atoi(os.Getenv("DB_POOL_MAX"))
	dbConfiguration.databaseMaxIdle, err = strconv.Atoi(os.Getenv("DB_POOL_MAX_IDLE"))
	dbConfiguration.logging, err = strconv.ParseBool(os.Getenv("DB_LOGGING"))
	if err != nil {
		panic("Error parsing the database port")
	}
	dbConfiguration.databaseName = os.Getenv("DB_NAME")
}

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
