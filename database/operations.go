package database

import "ascan/desafio-go/model"

func (db *database) CreateUser(user *model.User) (*model.User, error) {
	result := db.db.Create(user)
	return user, result.Error
}

func (db *database) CreateUsersWithArray(users []*model.User) ([]*model.User, error) {
	result := db.db.Create(&users)
	return users, result.Error
}

func (db *database) GetUserByUsername(username string) (*model.User, error) {
	user := &model.User{}
	result := db.db.Where("username = ?", username).First(user)
	return user, result.Error
}

func (db *database) EditUserByUsername(newUser *model.User, username string) (*model.User, error) {
	user := &model.User{}
	result := db.db.Where("username = ?", username).First(user)

	if result.Error != nil {
		return nil, result.Error
	}
	
	user.Username = newUser.Username
	user.FirstName = newUser.FirstName
	user.LastName = newUser.LastName
	user.Email = newUser.Email
	user.Password = newUser.Password
	user.Phone = newUser.Phone

	result = db.db.Save(user)

	return user, result.Error
}

func (db *database) DeleteUserByUsername(username string) (*model.User, error) {
	user := &model.User{}
	result := db.db.Where("username = ?", username).Delete(user)
	return user, result.Error
}