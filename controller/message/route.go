package message

import "github.com/gin-gonic/gin"

func InitRoute(engine *gin.Engine) {
	engine.GET("/messages", Message)
	return
}
