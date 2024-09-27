package route

import (
	"board-api-server/api/controller"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Engine         *gin.Engine
	UserController *controller.UserController
}

func (c *Config) Setup() {
	api := c.Engine.Group("/api")
	api.POST("/sign-up", c.UserController.SignUp)
}
