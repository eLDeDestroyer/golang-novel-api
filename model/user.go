package model

type User struct {
	Id        int    `json:"id" gorm:"primary_key;column:id"`
	Username  string `json:"username" gorm:"column:username"`
	Name      string `json:"name" gorm:"column:name"`
	Password  string `json:"password" gorm:"column:password"`
	IsPremium bool   `json:"is_premium" gorm:"column:is_premium"`
}
