package entities

import "time"

type SQLModelInit struct {
	Id        int        `json:"id" gorm:"column:id;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

type SQLModelUpdate struct {
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
