package dto

type RequestPage struct {
	BookId int `gorm:"column:book_id" json:"book_id"`
	Page []int `gorm:"column:page" json:"page"`
}