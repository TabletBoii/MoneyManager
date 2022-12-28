package models

import "time"

type Turnover struct {
	UserID        int       `gorm:"foreign_key" json:"user_id"`
	CategoryID    int       `gorm:"foreign_key" json:"category_id"`
	TurnoverName  string    `json:"turnover_name"`
	TurnoverType  string    `json:"turnover_type"`
	TurnoverValue int       `json:"turnover_value"`
	TurnoverDate  time.Time `json:"turnover_date"`
}

func (Turnover) TableName() string {
	return "user_turnover"
}
