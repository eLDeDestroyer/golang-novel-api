package dto

type RequestAction struct {
	UserId int `json:"user_id" gorm:"column:user_id"`
	BookId int `json:"book_id" gorm:"column:book_id"`
}