package robotmsg

import (
	"fmt"
	"github.com/grearter/rpa-agent/api"
	"github.com/sirupsen/logrus"
)

// List 获取未Pulled的消息列表
func List() (messageAPIs []*api.RobotMessage, err error) {
	sql := fmt.Sprintf("SELECT * from %s WHERE pulled = false", tableName)

	rows, err := sqliteDB.Query(sql)
	if err != nil {
		logrus.Error("sqlite query err: %s, sql: %s", err.Error(), sql)
		return
	}

	for rows.Next() {
		m := &api.RobotMessage{
			ID:      0,
			RobotID: "",
			Process: "",
			Level:   "",
			Ct:      0,
			Content: "",
			Pulled:  false,
		}

		if err := rows.Scan(&m.ID, &m.RobotID, &m.Process, &m.Level, &m.Ct, &m.Content, &m.Pulled); err != nil {
			logrus.Errorf("rows.Scan err: %s", err.Error())
			return nil, err
		}

		messageAPIs = append(messageAPIs, m)
	}

	return
}
