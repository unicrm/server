package models

import "gorm.io/gorm"

type Test struct {
	gorm.Model
}

func (Test) TableName() string {
	return "tests"
}
