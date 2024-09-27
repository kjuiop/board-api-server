package route

import "github.com/gin-gonic/gin"

type Config struct {
	Engine *gin.Engine
}

func (c *Config) Setup() {
}
