package dto

type RequestBookCategory struct {
	BookId     int   `gorm:"column:book_id" json:"book_id"`
	CategoryId []int `json:"category_id"`
}
