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
	c.SetupMember()
}

func (c *Config) SetupMember() {
	member := c.Engine.Group("/member")
	member.POST("/sign-up", c.MemberController.SignUp)
}
