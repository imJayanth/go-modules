package config

import (
	"fmt"
	"log"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (appConfig *AppConfig) SetupDatabase() {
	gormConfig := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	}
	db, e := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@(%s:%s)/?charset=utf8&loc=%s&parseTime=True",
		appConfig.MySQLDBConfig.DBUser,
		appConfig.MySQLDBConfig.DBPassword,
		appConfig.MySQLDBConfig.DBHost,
		appConfig.MySQLDBConfig.DBPort,
		url.QueryEscape(appConfig.ServerConfig.Timezone),
	)), gormConfig)
	if e != nil {
		log.Fatalf("Error Connecting to the DB: %s", e.Error())
	}
	if e = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v", appConfig.MySQLDBConfig.DBName)).Error; e != nil {
		log.Fatalf("failed to create database: %v\n", e.Error())
	}
	if appConfig.MySQLDBConfig.Log {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}
	if db, e = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&loc=%s&parseTime=True",
		appConfig.MySQLDBConfig.DBUser,
		appConfig.MySQLDBConfig.DBPassword,
		appConfig.MySQLDBConfig.DBHost,
		appConfig.MySQLDBConfig.DBPort,
		appConfig.MySQLDBConfig.DBName,
		url.QueryEscape(appConfig.ServerConfig.Timezone),
	)), gormConfig); e != nil {
		log.Fatalf("Error Connecting to the DB: %s", e.Error())
	}
	appConfig.MySQLDBConfig.DB = db
	sqldb, _ := db.DB()
	if e := sqldb.Ping(); e != nil {
		log.Fatalf("Error while pinging the DB: %s", e.Error())
	}
	for _, model := range appConfig.MySQLDBConfig.Automigrate {
		if e := appConfig.MySQLDBConfig.DB.Table(model.TableName()).AutoMigrate(model); e != nil {
			log.Fatalf("Error automigrating %s: %s", model.TableName(), e.Error())
		}
	}
	log.Printf("Connected to the DB: %s\n", appConfig.MySQLDBConfig.DBName)
}
