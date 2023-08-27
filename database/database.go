package database

import (
	"fmt"

	"example.com/todo-app/models"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Init() {
	username := viper.GetString("USERNAME")
	password := viper.GetString("PASSWORD")
	port := viper.GetString("DB_PORT")
	dbname := viper.GetString("DB_NAME")

	dsn := username + ":" + password + "@tcp(0.0.0.0:" + port + ")/" + dbname
	var err error
	Database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// if database do not exist, create it
		fmt.Println("Trying to create database : " + dbname + "...")
		dsn = username + ":" + password + "@tcp(0.0.0.0:" + port + ")/"
		Database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Println(err.Error())
		}
		Database.Exec("CREATE DATABASE " + dbname)
		Database.Exec("USE " + dbname)
		fmt.Println("Database created : " + dbname)
	} else {
		fmt.Println("Connected to database : " + dbname)
	}
	Migrate()
}

func Migrate() {
	err := Database.AutoMigrate(&models.Todo{})
	if err != nil {
		fmt.Println(err.Error())
	}
}
