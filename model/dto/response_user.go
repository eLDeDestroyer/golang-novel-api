package dto

type ResponseUser struct {
	Username string             `gorm:"column:username" json:"username"`
	Name     string             `grom:"column:name" json:"name"`
	Books    []BookResponseUser `json:"books"`
}

type BookResponseUser struct {
	Id          int    `gorm:"primary_key;column:id" json:"id"`
	Title       string `grom:"column:title" json:"title"`
	ImagePath   string `grom:"column:image_path" json:"image_path"`
	Description string `grom:"column:description" json:"description"`
	PageCount   int    `grom:"column:page_count" json:"page_count"`
}
