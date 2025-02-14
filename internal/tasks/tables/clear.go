package tables

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func ClearTable(db *gorm.DB, tableName string) error {
	if tableName == "" {
		return errors.New("表名不能为空")
	}
	detail := &ClearDB{
		TableName:    tableName,
		CompareField: "created_at",
		Interval:     "10m",
	}
	if db == nil {
		return errors.New("数据库未初始化")
	}
	duration, err := time.ParseDuration(detail.Interval)
	if err != nil {
		return err
	}
	if duration < 0 {
		return errors.New("时间间隔不能为负数")
	}
	err = db.Debug().Exec(fmt.Sprintf("DELETE FROM %s WHERE %s < ?", detail.TableName, detail.CompareField), time.Now().Add(-duration)).Error
	if err != nil {
		return err
	}
	return nil
}
