package model

import (
	"gogo/common"
)

type Menu struct {
	common.SQLModelInit
	Name string `json:"name" form:"name" gorm:"column:name;"`
}

func (Menu) TableName() string {
	return "menus"
}

type MenuCreation struct {
	Id   int     `json:"id" gorm:"column:id;"`
	Name *string `json:"name" form:"name" gorm:"column:name;"`
}

func (MenuCreation) TableName() string {
	return Menu{}.TableName()
}

type MenuUpdate struct {
	common.SQLModelUpdate
	Name string `json:"name" form:"name" gorm:"column:name;"`
}

func (MenuUpdate) TableName() string {
	return Menu{}.TableName()
}
