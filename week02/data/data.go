package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Data struct {
	// TODO wrapped database client
	db  *gorm.DB

}

func NewData() *Data{
	dsn := "root:mysql@tcp(127.0.0.1:3306)/test?charset=utf8"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic("connect db error")
	}

	return &Data{
		db: db,
	}
}
