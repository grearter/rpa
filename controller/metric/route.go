package metric

import "github.com/gin-gonic/gin"

func InitRoute(engine *gin.Engine) {
	engine.GET("/host_metrics", HostMetric)
	engine.GET("/robot_metrics", RobotMetric)

	return
}
