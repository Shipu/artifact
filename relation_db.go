package artifact

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *Database

type Database struct {
	*gorm.DB
}

func NewDatabase() *Database {
	dsn := Config.GetString(Config.RelationDBConfig+".Username") + ":" + Config.GetString(Config.RelationDBConfig+".Password") + "@tcp(" + Config.GetString(Config.RelationDBConfig+".Host") + ":" + Config.GetString(Config.RelationDBConfig+".Port") + ")/" + Config.GetString(Config.RelationDBConfig+".Database") + "?charset=utf8mb4&parseTime=True&loc=Local"
	connection := Config.GetString(Config.RelationDBConfig + ".Connection")
	err := error(nil)
	var db *gorm.DB
	if connection == "mysql" {
		dsn = Config.GetString(Config.RelationDBConfig+".Username") + ":" + Config.GetString(Config.RelationDBConfig+".Password") + "@tcp(" + Config.GetString(Config.RelationDBConfig+".Host") + ":" + Config.GetString(Config.RelationDBConfig+".Port") + ")/" + Config.GetString(Config.RelationDBConfig+".Database") + "?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else if connection == "sqlite" {
		dsn = Config.GetString(Config.RelationDBConfig + ".Database")
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	} else if connection == "postgres" {
		dsn = "host=" + Config.GetString(Config.RelationDBConfig+".Host") + " user=" + Config.GetString(Config.RelationDBConfig+".Username") + " password=" + Config.GetString(Config.RelationDBConfig+".Password") + " dbname=" + Config.GetString(Config.RelationDBConfig+".Database") + " port=" + Config.GetString(Config.RelationDBConfig+".Port") + " sslmode=disable TimeZone=Asia/Shanghai"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		panic(err)
	}

	return &Database{db}
}
