package entities

type Food struct {
	SQLModelInit
	Name       string `json:"name" form:"name" gorm:"column:name;"`
	Price      int64  `json:"price" form:"price" gorm:"column:price;"`
	FoodImage  string `json:"food_image" form:"food_image" gorm:"column:food_image;"`
	MenuId     int    `json:"menu_id" form:"menu_id" gorm:"column:menu_id;"`
	CategoryId int    `json:"category_id" form:"category_id" gorm:"column:category_id;"`
}
