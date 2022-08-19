package rest

import (
	"ascan/desafio-go/model"

	"github.com/gin-gonic/gin"
)

func (r rest) create(c *gin.Context) {
	user := new(model.User)
	c.BindJSON(user)

	user, err := r.service.CreateUser(user)

	r.returnDataOrOk(user, err, c)
}

func (r rest) createWithArray(c *gin.Context) {
	users := make([]*model.User, 0)
	c.BindJSON(&users)

	users, err := r.service.CreateUsersWithArray(users)

	r.returnDataOrOk(users, err, c)
}

func (r rest) getUserByUsername(c *gin.Context) {
	username := c.Param("username")

	user, err := r.service.GetUserByUsername(username)

	r.returnDataOrOk(user, err, c)
}

func (r rest) editUserByUsername(c *gin.Context) {
	username := c.Param("username")

	newUser := new(model.User)
	c.BindJSON(newUser)

	user, err := r.service.EditUserByUsername(newUser, username)

	r.returnDataOrOk(user, err, c)
}

func (r rest) deleteUserByUsername(c *gin.Context) {
	username := c.Param("username")

	user, err := r.service.DeleteUserByUsername(username)

	r.returnDataOrOk(user, err, c)
}
