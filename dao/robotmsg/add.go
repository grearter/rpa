package robotmsg

import (
	"fmt"
	"github.com/grearter/rpa-agent/api"
	"github.com/sirupsen/logrus"
	"reflect"
	"strings"
)

// Add 新增RobotMessage
func Add(messageAPI *api.RobotMessage) (err error) {
	t := reflect.TypeOf(*messageAPI)
	v := reflect.ValueOf(*messageAPI)

	fieldNum := t.NumField()
	fields := make([]string, 0, fieldNum)
	values := make([]string, 0, fieldNum)
	execArgs := make([]interface{}, 0, fieldNum)

	sql := fmt.Sprintf("INSERT INTO %s", tableName)

	for i := 0; i < fieldNum; i++ {
		field := t.Field(i)
		sqliteTagValues := strings.Split(field.Tag.Get("sqlite"), ",")

		if len(sqliteTagValues) < 2 {
			return fmt.Errorf("fieldName '%s' sqlite tag err", field.Name)
		}

		sqlFieldName := sqliteTagValues[0]

		if sqlFieldName == "id" && messageAPI.ID <= 0 {
			continue
		}

		fields = append(fields, sqlFieldName)
		values = append(values, "?")
		execArgs = append(execArgs, v.Field(i).Interface())
	}

	sql += fmt.Sprintf(" (%s) values (%s) ", strings.Join(fields, ","), strings.Join(values, ","))

	logrus.Infof("msg: %+v, sql: %s", messageAPI, sql)

	stmt, err := sqliteDB.Prepare(sql)
	if err != nil {
		logrus.Errorf("sqlite prepare err: %s, sql: %s", err.Error(), sql)
		return
	}

	logrus.Infof("exec args: %v", execArgs)

	if _, err = stmt.Exec(execArgs...); err != nil {
		logrus.Errorf("sqlite exec err: %s, args: %v", err.Error(), execArgs)
		return
	}

	return
}
