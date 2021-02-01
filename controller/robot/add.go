package robot

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/grearter/rpa-agent/util"
	"github.com/sirupsen/logrus"
	"net/http"
)

type startReq struct {
	RobotID  string `json:"robotId" binding:"required"`
	Filepath string `json:"filepath" binding:"required"`
}

func (req *startReq) regular() error {
	if req.RobotID == "" {
		return errors.New("robotId为空")
	}

	if req.Filepath == "" {
		return errors.New("filepath为空")
	}

	return nil
}

// Add 添加/启动机器人
func Add(c *gin.Context) {
	req := new(startReq)

	if err := c.ShouldBindBodyWith(req, binding.JSON); err != nil {
		logrus.Errorf("parse param err: %s, req: %+v", err.Error(), req)
		c.JSON(http.StatusBadRequest, &util.Resp{Code: util.CodeParamErr, Msg: err.Error()})
		return
	}

	// TODO: 判断机器人状态

	logrus.Infof("start robot success")

	c.JSON(http.StatusCreated, &util.Resp{Code: util.CodeOK})
	return
}
