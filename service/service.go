package service

import (
	"ascan/desafio-go/database"
	"ascan/desafio-go/model"
	"errors"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type Service interface {
	CreateUser(user *model.User) (*model.User, error)
	CreateUsersWithArray(users []*model.User) ([]*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	EditUserByUsername(newUser *model.User, username string) (*model.User, error)
	DeleteUserByUsername(username string) (*model.User, error)
}

type service struct {
	database database.Database
}

func NewService(database database.Database) Service {
	svc := &service{
		database: database,
	}

	return svc
}

func (svc *service) CreateUser(user *model.User) (*model.User, error) {
	if user.Username == "" || 
	user.Password == "" || 
	user.Email == "" || 
	user.FirstName == "" || 
	user.LastName == "" ||  
	user.Phone == "" {
		return nil, errors.New("invalid user data")
	}

	log.Infof("Creating user %s", user.Username)

	user.ID = uuid.New().String()
	user.UserStatus = 0

	return svc.database.CreateUser(user)
}

func (svc *service) CreateUsersWithArray(users []*model.User) ([]*model.User, error) {
	if len(users) == 0 {
		return nil, errors.New("invalid user data")
	}

	log.Infof("Creating users with array")

	return svc.database.CreateUsersWithArray(users)
}

func (svc *service) GetUserByUsername(username string) (*model.User, error) {
	log.Infof("Getting user %s", username)

	return svc.database.GetUserByUsername(username)
}

func (svc *service) EditUserByUsername(updatedUser *model.User, username string) (*model.User, error) {
	log.Infof("Editing user %s", username)

	return svc.database.EditUserByUsername(updatedUser, username)
}

func (svc *service) DeleteUserByUsername(username string) (*model.User, error) {
	log.Infof("Deleting user %s", username)

	return svc.database.DeleteUserByUsername(username)
}