package metric

import (
	"github.com/gin-gonic/gin"
	"github.com/grearter/rpa-agent/api"
	"github.com/grearter/rpa-agent/util"
	"github.com/sirupsen/logrus"
	"net/http"
)

// RobotMetric 机器人指标
func RobotMetric(c *gin.Context) {

	logrus.Infof("get robot metric")

	c.JSON(http.StatusOK, util.NewRespWithData([]*api.Robot{}))
	return
}
