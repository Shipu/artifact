package artifact

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *Database

type Database struct {
	*gorm.DB
}

func NewDatabase() *Database {
	dsn := Config.GetString("DB.Username")+":"+Config.GetString("DB.Password")+"@tcp("+Config.GetString("DB.Host")+":"+Config.GetString("DB.Port")+")/"+Config.GetString("DB.Database")+"?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "root:root@tcp(localhost:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return &Database{db}
}