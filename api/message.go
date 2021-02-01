package api

type RobotMessage struct {
	ID      int    `json:"id" sqlite:"id,INTEGER PRIMARY KEY autoincrement"` // 消息ID, 自动递增
	RobotID string `json:"robotId" sqlite:"robotId,VARCHAR(64)"`             // 机器人ID
	Process string `json:"process" sqlite:"process,VARCHAR(64)"`             // 机器人流程名称
	Level   string `json:"level" sqlite:"level,VARCHAR(16)"`                 // 日志级别
	Ct      int64  `json:"ct" sqlite:"ct,INTEGER"`                           // 日志时间
	Content string `json:"content" sqlite:"content,TEXT"`                    // 日志描述
	Pulled  bool   `json:"-" sqlite:"pulled,BOOLEAN"`                        // server端是否已拉取此条msg
}
