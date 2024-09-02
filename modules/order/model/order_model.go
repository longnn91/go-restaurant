package model

import (
	"gogo/common"
)

type Order struct {
	common.SQLModelInit
	UserId   int    `json:"user_id" form:"user_id" gorm:"column:user_id;"`
	UserNote string `json:"user_note" form:"user_note" gorm:"column:user_note;"`
}

func (Order) TableName() string {
	return "orders"
}

type OrderCreation struct {
	Id       int    `json:"id" gorm:"column:id;"`
	UserId   int    `json:"user_id" form:"user_id" gorm:"column:user_id;"`
	UserNote string `json:"user_note" form:"user_note" gorm:"column:user_note;"`
}

func (OrderCreation) TableName() string {
	return Order{}.TableName()
}

type OrderUpdate struct {
	common.SQLModelUpdate
	UserNote string `json:"user_note" form:"user_note" gorm:"column:user_note;"`
}

func (OrderUpdate) TableName() string {
	return Order{}.TableName()
}
