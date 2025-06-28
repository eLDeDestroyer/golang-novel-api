package model

type Like struct {
	Id     int `gorm:"column:id; primary_key" json:"id"`
	UserId int `gorm:"column:user_id" json:"user_id"`
	BookId int `gorm:"column:book_id" json:"book_id"`
}
