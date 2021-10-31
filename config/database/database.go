package database

import (
	"fmt"
	env "github.com/nogopy/jwt-for-authentication/config/environment"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase() *gorm.DB {
	var db *gorm.DB

	USER := env.GetConfiguration().Database.DbUser
	PASS := env.GetConfiguration().Database.DbPassword
	HOST := env.GetConfiguration().Database.DbHost
	DBNAME := env.GetConfiguration().Database.DbName

	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, DBNAME)
	//fmt.Println(URL)
	db, err := gorm.Open(mysql.Open(URL), GetGormConfig())

	if err != nil {
		log.Errorf("Failed to connect to Mysql database, exception %s", err.Error())
		panic("Failed to connect to Mysql database")

	}
	fmt.Println("Mysql connection established")

	return db
}

func GetGormConfig() *gorm.Config {
	logMode := logger.Silent
	if env.GetConfiguration().AppEnv == "LOCAL" {
		logMode = logger.Info
	}

	return &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	}
}
