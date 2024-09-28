package route

import (
	"board-api-server/api/controller"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Engine           *gin.Engine
	MemberController *controller.MemberController
}

func (c *Config) Setup() {
	api := c.Engine.Group("/api")
	api.POST("/sign-up", c.MemberController.SignUp)
}
