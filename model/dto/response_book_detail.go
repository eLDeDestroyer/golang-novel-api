package dto

type ResponseBookDetail struct {
	Name        string                    `gorm:"column:name" json:"name"`
	Title       string                    `gorm:"column:title" json:"title"`
	Description string                    `gorm:"column:description" json:"description"`
	ImagePath   string                    `gorm:"column:image_path" json:"image_path"`
	Action      ActionResponseBookDetail  `json:"action"`
	Pages       []PagesResponseBookDetail `json:"pages"`
}

type ActionResponseBookDetail struct {
	Seen int `gorm:"column:seen" json:"seen"`
	Like int `gorm:"column:like" json:"like"`
	Page int `gorm:"column:page" json:"page"`
}

type PagesResponseBookDetail struct {
	Id   int    `gorm:"primary_key;column:id" json:"id"`
	Page int    `gorm:"column:page" json:"page"`
	Text string `gorm:"column:text" json:"text"`
}
