package metric

import (
	"github.com/gin-gonic/gin"
	"github.com/grearter/rpa-agent/api"
	"github.com/grearter/rpa-agent/util"
	"net/http"
)

// HostMetric 主机指标
func HostMetric(c *gin.Context) {

	cpuInfo := util.GetCpuInfo()
	memInfo := util.GetMemInfo()

	hostMetric := &api.HostMetric{
		Cpu: cpuInfo,
		Mem: memInfo,
	}

	c.JSON(http.StatusOK, util.NewRespWithData(hostMetric))
	return
}
