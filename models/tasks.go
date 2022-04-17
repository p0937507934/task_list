package models

type Tasks struct {
	Id     int    `gorm:"column:id"`
	Name   string `gorm:"column:name;comment:'任務名'"`
	Status int    `gorm:"column:status;comment:'0: 未完成, 1: 已完成'"`
}

func (*Tasks) TableName() string {
	return "task"
}
