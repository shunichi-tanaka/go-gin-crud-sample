package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/v1")
	{
		apiv1.GET("/user", GetUser)
	}

	return r
}

type User struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
}

func GetUser(c *gin.Context) {
	c.JSON(200, User{"1", "user_name"})
}
