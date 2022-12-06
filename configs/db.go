package config

import (
	"log"
	"os"
	"sync"
	"time"

	lModel "go-blogs-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBO *gorm.DB
var initOnce sync.Once
var migrateOnce sync.Once
var connErr error

func InitDB() {
	connectStr := os.Getenv("CONNECTION_STRING")
	dbLog := os.Getenv("DB_LOG")
	gormConfig := new(gorm.Config)

	oneTimeDBConnSetup := func() {
		DBO, connErr = gorm.Open(mysql.Open(connectStr), gormConfig)
	}

	if dbLog == "Y" {
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Disable color
			},
		)

		gormConfig.Logger = newLogger

	}

	initOnce.Do(oneTimeDBConnSetup)

	if connErr != nil {
		panic(connErr)
	}

	migrateOnce.Do(migrateTables)
}

func migrateTables() {
	DBO.AutoMigrate(&lModel.Article{})
	DBO.AutoMigrate(&lModel.Comment{})
}
