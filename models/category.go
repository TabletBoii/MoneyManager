package models

type Category struct {
	CategoryID   int    `gorm:"primary_key" json:"category_id"`
	CategoryName string `json:"category_name"`
}

func (Category) TableName() string {
	return "categories"
}

// func NewCategory

// func () GetCategory() {

// }
