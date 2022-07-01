package config

import (
	"database/sql"
	"errors"
	"fmt"
	util "livecode-gorm-wmb/utils"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dBConfig struct {
	dbVenv string
	dbHost string
	dbPort string
	dbUser string
	dbPass string
	dbName string
}

type Config struct {
	db *gorm.DB
}

func (c *Config) initDb() {
	dbVenv := os.Getenv("DB_ENV")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dBConfig := dBConfig{
		dbVenv, dbHost, dbPort, dbUser, dbPass, dbName,
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dBConfig.dbHost, dBConfig.dbUser, dBConfig.dbPass, dBConfig.dbName, dBConfig.dbPort)

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	util.PanicError(err)

	c.checkConnection(gormDB)

	if dBConfig.dbVenv == "development" {
		log.Println("Running in Development Environment")
		c.db = gormDB.Debug()
	} else if dBConfig.dbVenv == "production" {
		log.Println("Running in Production Environment")
		c.db = gormDB
	} else {
		util.PanicError(errors.New(fmt.Sprintf("%s in wrong environment",dBConfig.dbVenv)))
	}
}

func (c *Config) checkConnection (gormDB *gorm.DB) {
	runningDb, err := gormDB.DB()
	util.PanicError(err)
	err = runningDb.Ping()
	util.PanicError(err)
}

func (c *Config) DbConn() *gorm.DB {
	return c.db
}

func (c *Config) DbClose() {
	runningDb, err := c.db.DB()
	util.PanicError(err)
	defer func(runningDb *sql.DB) {
		err := runningDb.Close()
		util.PanicError(err)
	}(runningDb)
}
func NewConfigDB() *Config {
	cfg := Config{}
	cfg.initDb()
	return &cfg
}
