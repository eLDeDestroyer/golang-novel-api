package model

type Category struct {
	Id     int `gorm:"column:id; primary_key" json:"id"`
	Category string `gorm:"column:category" json:"category"`
}