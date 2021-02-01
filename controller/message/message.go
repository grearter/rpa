package message

import (
	"github.com/gin-gonic/gin"
	"github.com/grearter/rpa-agent/dao/robotmsg"
	"github.com/grearter/rpa-agent/util"
	"net/http"
)

// Message 拉取Message
func Message(c *gin.Context) {
	messageAPIs, err := robotmsg.List()
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code": 1,
			"msg":  err.Error(),
		})
	}

	c.JSON(http.StatusOK, &util.Resp{Code: util.CodeOK, Data: messageAPIs})
	return
}
