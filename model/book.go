package model

type Book struct {
	Id    int    `gorm:"primary_key; column:id" json:"id"`
	Title string `gorm:"column:title" json:"title"`
	Description string `gorm:"column:description" json:"description"`
	ImagePath string `gorm:"column:image_path" json:"image_path"`
	UserId int `gorm:"column:user_id" json:"user_id"`
}
