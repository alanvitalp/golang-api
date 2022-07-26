package rest

import (
	"ascan/desafio-go/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
	"gorm.io/gorm"
)

type Rest interface {
	ListenAndServe() error
}

type rest struct {
	httpPort 	int
	service 	service.Service
}

func NewRest(service service.Service) Rest {
	config := configFromEnv()

	return rest{
		httpPort: config.port,
		service:  service,
	}
}

func (r rest) ListenAndServe() error {
	ginEngine := gin.New()
	ginEngine.Use(ginlogrus.Logger(log.New()), gin.Recovery())

	ginEngine.POST("user", r.create)
	ginEngine.POST("user/createWithArray", r.createWithArray)
	ginEngine.GET("user/:username", r.getUserByUsername)
	ginEngine.PUT("user/:username", r.editUserByUsername)
	ginEngine.DELETE("user/:username", r.deleteUserByUsername)

	return ginEngine.Run(fmt.Sprintf("0.0.0.0:%d", r.httpPort))
}

func (r rest) returnDataOrOk(data interface{}, err error, c *gin.Context) {
	if err != nil {
		log.Errorf("Failed to get data %s, caused by: %v", c.Request.URL, err.Error())

		status := http.StatusInternalServerError
		if err == gorm.ErrRecordNotFound {
			status = http.StatusNotFound
		}

		c.JSON(status, gin.H{
			"message": err.Error(),
		})
	} else {
		if data != nil{
			c.JSON(http.StatusOK, data)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
			})
		}
	}
}