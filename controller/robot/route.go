package robot

import "github.com/gin-gonic/gin"

func InitRoute(engine *gin.Engine) {
	engine.POST("/robots", Add)
	engine.DELETE("/robots/:robotId", Delete)

	return
}
