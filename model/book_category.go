package model

type BookCategory struct {
	Id         int `gorm:"column:id; primary_key" json:"id"`
	CategoryId int `gorm:"column:category_id" json:"category_id"`
	BookId     int `gorm:"column:book_id" json:"book_id"`
}
