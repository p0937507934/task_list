package middleware

import (
	"encoding/json"
	"fmt"
	"strings"
	"tasks_list/driver"
	custom_error "tasks_list/pkg/error"

	"github.com/gin-gonic/gin"
)

func HandleError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		errs := c.Errors
		if len(errs) > 0 {
			for _, err := range errs {
				switch err.Err.(type) {
				case *custom_error.CustomError:
					customErr := err.Err.(*custom_error.CustomError)
					v, marshalErr := json.Marshal(customErr.Value)
					if marshalErr != nil {
						panic(marshalErr)
					}
					s := strings.Builder{}
					s.Write([]byte(fmt.Sprintf("StatusCode: %d,PrivateError: %s, value: %s", customErr.StatusCode, customErr.PrivateMsg, string(v))))
					if customErr.Err != nil {
						s.Write([]byte(fmt.Sprintf(", Error: %s", customErr.Err.Error())))
					}
					driver.Logger.Error(s.String())
					c.AbortWithStatusJSON(customErr.StatusCode, customErr.NewResponseMsg())
					return

				default:
					c.AbortWithStatusJSON(400, err)
					return
				}
			}
		}
	}
}
