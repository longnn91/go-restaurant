package entities

type User struct {
	SQLModelInit
	Username    string `json:"username" form:"username" gorm:"column:username;not null;unique"`
	Password    string `json:"password" form:"password" gorm:"column:password;not null"`
	FirstName   string `json:"first_name" form:"first_name" gorm:"column:first_name"`
	LastName    string `json:"last_name" form:"last_name" gorm:"column:last_name"`
	PhoneNumber string `json:"phone_number" form:"phone_number" gorm:"column:phone_number"`
}

type UserLogin struct {
	Username string `json:"username" form:"username" gorm:"column:username;not null;unique"`
	Password string `json:"password" form:"password" gorm:"column:password;not null"`
}
