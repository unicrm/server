package tables

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func ClearTables(db *gorm.DB) error {
	if db == nil {
		return errors.New("数据库未初始化")
	}
	var clearTablesDetail []ClearDB
	// 清理jwt黑名单表，保留7天数据
	clearTablesDetail = append(clearTablesDetail, ClearDB{
		TableName:    "unicrm_jwt_black_list",
		CompareField: "created_at",
		Interval:     "168h",
	})
	for _, detail := range clearTablesDetail {
		err := ClearTable(db, &detail)
		if err != nil {
			return err
		}
	}
	return nil
}

func ClearTable(db *gorm.DB, detail *ClearDB) error {
	if detail.TableName == "" {
		return errors.New("表名不能为空")
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
