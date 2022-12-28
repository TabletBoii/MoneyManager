package models

// "github.com/jinzhu/gorm"

type User struct {
	UserID     int    `gorm:"primary_key" json:"user_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Password   string `json:"password"`
	UserStatus string `json:"user_status"`
}

func (User) TableName() string {
	return "users"
}

// func GetUser() (*User, error) {
// 	var user User
// 	err := db.Where

// 	return user, nil
// }
