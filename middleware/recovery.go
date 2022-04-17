package middleware

import (
	"tasks_list/driver"

	"github.com/gin-gonic/gin"
)

func GinRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				driver.Logger.Panic(err)
			}
		}()
	}
}
