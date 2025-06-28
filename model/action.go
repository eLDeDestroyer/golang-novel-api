package model

type Action struct {
	Id     int `gorm:"primary_key; column:id" json:"id"`
	UserId int `json:"user_id" gorm:"column:user_id"`
	BookId int `json:"book_id" gorm:"column:book_id"`
}
