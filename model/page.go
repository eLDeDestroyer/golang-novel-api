package model

type Page struct {
	Id    int    `gorm:"primary_key; column:id" json:"id"`
	Page int `gorm:"column:page" json:"page"`
	Text string `gorm:"column:text" json:"text"`
	BookId int `gorm:"column:book_id" json:"book_id"`
}