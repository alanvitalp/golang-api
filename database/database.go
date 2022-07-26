package database

import (
	"ascan/desafio-go/model"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database interface {
	CreateUser(user *model.User) (*model.User, error)
	CreateUsersWithArray(users []*model.User) ([]*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	EditUserByUsername(newUser *model.User, username string) (*model.User, error)
	DeleteUserByUsername(username string) (*model.User, error)
}

type database struct {
	db *gorm.DB
}

const (
	maxIdleTime 				= 	time.Minute
	maxIdleConnections 	= 	10
	maxOpenConnections 	= 	10
)

func NewDatabase() (Database, error) {
	config := configFromEnv()

	db, err := gorm.Open(
		postgres.Open(config.dsn()),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		},
	)

	if err != nil {
		return nil, err
	}

	sqlDb, err := db.DB()

	if err != nil {
		return nil, err
	}

	sqlDb.SetConnMaxIdleTime(maxIdleTime)
	sqlDb.SetMaxIdleConns(maxIdleConnections)
	sqlDb.SetMaxOpenConns(maxOpenConnections)

	if err := db.AutoMigrate(&model.User{}); err != nil {
		return nil, err
	}

	return &database{db: db}, nil
}