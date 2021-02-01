package robot

import (
	"github.com/gin-gonic/gin"
	"github.com/grearter/rpa-agent/util"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Delete 删除/停止机器人
func Delete(c *gin.Context) {
	robotID := c.Param("robotId")

	logrus.Infof("stop robot '%s' success", robotID)

	c.JSON(http.StatusOK, util.NewRespWithData(nil))
	return
}
