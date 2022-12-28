package lib

import "gorm.io/gorm"

// import (
// 	"fmt"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

type Database struct {
	*gorm.DB
}

func NewDatabase(env Env, logger Logger) Database {

}
