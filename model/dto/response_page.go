package dto

type ResponsePage struct {
	Text string `gorm:"column:text" json:"text"`
	Page int `gorm:"column:page" json:"page"`
	PrevPage *int `gorm:"column:prev_page" json:"prev_page"`
	NextPage *int `gorm:"column:next_page" json:"next_page"`
}