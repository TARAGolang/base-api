package middlewares

import (
	"github.com/pushpaldev/base-api/config"
	"github.com/spf13/viper"
	"gopkg.in/gin-gonic/gin.v1"
)

func ConfigMiddleware(viper *viper.Viper) gin.HandlerFunc {
	return func(c *gin.Context) {
		config.ToContext(c, config.New(viper))
		c.Next()
	}
}
