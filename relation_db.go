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
	dsn := Config.GetString(Config.RelationDBConfig+"DB.Username") + ":" + Config.GetString(Config.RelationDBConfig+".Password") + "@tcp(" + Config.GetString(Config.RelationDBConfig+".Host") + ":" + Config.GetString(Config.RelationDBConfig+".Port") + ")/" + Config.GetString(Config.RelationDBConfig+".Database") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return &Database{db}
}
