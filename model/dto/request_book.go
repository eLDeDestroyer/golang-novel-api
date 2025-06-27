package dto

type RequestBook struct {
	Title string `gorm:"column:title" json:"title" form:"title"`
	Description string `gorm:"column:description" json:"description" form:"description"`
}