package router

import (
	"sms/config"
	"sms/routes"

	"github.com/gin-gonic/gin"
)

func Init() {
	route := gin.Default()
	routes.PingRoutes(route)
	routes.AuthRoutes(route)
	routes.AttendanceRoutes(route)
	routes.StudentRoutes(route)
	routes.WorkDoneRoutes(route)
	route.Run(config.Conf.Port)
}
