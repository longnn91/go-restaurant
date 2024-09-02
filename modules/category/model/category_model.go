package model

import (
	"gogo/common"
)

type Category struct {
	common.SQLModelInit
	Name   string `json:"name" form:"name" gorm:"column:name;"`
	MenuId int    `json:"menu_id" form:"menu_id" gorm:"column:menu_id;"`
}

func (Category) TableName() string {
	return "categories"
}

type CategoryCreation struct {
	Id     int     `json:"id" gorm:"column:id;"`
	Name   *string `json:"name" form:"name" gorm:"column:name;"`
	MenuId *int    `json:"menu_id" form:"menu_id" gorm:"column:menu_id;"`
}

func (CategoryCreation) TableName() string {
	return Category{}.TableName()
}

type CategoryUpdate struct {
	common.SQLModelUpdate
	Name string `json:"name" form:"name" gorm:"column:name;"`
}

func (CategoryUpdate) TableName() string {
	return Category{}.TableName()
}
