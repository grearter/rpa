package robotmsg

import (
	"database/sql"
	"fmt"
	"github.com/grearter/rpa-agent/api"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"reflect"
	"strings"
)

var (
	sqliteDB  *sql.DB
	tableName = "robot_message"
)

// InitDB 初始化Table
func InitDB(dbFile string) (err error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		logrus.Errorf("open db file err: %s, file: %s", err.Error(), dbFile)
		return
	}
	sqliteDB = db

	t := reflect.TypeOf(api.RobotMessage{})
	filedNum := t.NumField()

	sqlStmt := fmt.Sprintf("CREATE TABLE %s (", tableName)
	for i := 0; i < filedNum; i++ {
		field := t.Field(i)
		sqliteTagValues := strings.Split(field.Tag.Get("sqlite"), ",")

		if len(sqliteTagValues) < 2 {
			return fmt.Errorf("fieldName '%s' sqlite tag err", field.Name)
		}

		sqlFieldName, sqlFieldType := sqliteTagValues[0], sqliteTagValues[1]

		sqlStmt += fmt.Sprintf("%s %s", sqlFieldName, sqlFieldType)

		if i < filedNum-1 {
			sqlStmt += ","
		}
	}

	sqlStmt += ")"

	fmt.Println(sqlStmt)

	if _, err := db.Exec(sqlStmt); err != nil {
		if err.Error() == "table robot_message already exists" {
			err = nil
		} else {
			logrus.Errorf("exec create table err: %s", err.Error())
			return err
		}
	}

	return
}
