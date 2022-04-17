package route

import (
	"tasks_list/config"
	"tasks_list/docs"
	"tasks_list/driver"
	"tasks_list/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init() {
	config.InitConfig()
	driver.InitGorm()
	driver.InitLogger()
}

func Serve() *gin.Engine {
	// init config and required
	Init()

	// init gin server
	r := gin.New()

	// init swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// init router
	r.Use(middleware.GinRecovery(), middleware.Logger(driver.Logger), middleware.HandleError())
	TaskRoute(r)

	return r
}
