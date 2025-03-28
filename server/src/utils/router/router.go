package router

import (
	"sms/config"
	"sms/middlewares"
	"sms/routes"

	"github.com/gin-gonic/gin"
)

func Init() {
	route := gin.Default()
	route.Use(middlewares.CORSMiddleware())
	routes.PingRoutes(route)
	routes.AuthRoutes(route)
	routes.AttendanceRoutes(route)
	routes.StudentRoutes(route)
	routes.WorkDoneRoutes(route)
	route.Run(config.Conf.Port)
}
