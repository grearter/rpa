package robotmsg

import (
	"encoding/json"
	"github.com/grearter/rpa-agent/api"
	"github.com/grearter/rpa-agent/dao/robotmsg"
	"github.com/sirupsen/logrus"
)

// HandlerRobotMessage 处理机器人消息
func HandlerRobotMessage(data []byte) {
	req := new(api.RobotMessage)

	if err := json.Unmarshal(data, req); err != nil {
		logrus.Errorf("parse req err: %s, req: %s", err.Error(), data)
		return
	}

	if err := robotmsg.Add(req); err != nil {
		logrus.Errorf("add robot msg err: %s, req: %+v", err.Error(), req)
		return
	}

	return
}
