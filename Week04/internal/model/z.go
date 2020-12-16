package model

// ZItem xxx
type ZItem struct {
	ID   int64  `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}
