package configs

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type config struct {
	User    string
	Passwd  string
	Addr    string
	Port    int
	DBName  string
	SSLMode string
}

var DB *gorm.DB

func InitDB() {
	var db = config{
		User:    viper.Get("DB_USERNAME").(string),
		Passwd:  viper.Get("DB_PASSWORD").(string),
		Addr:    viper.Get("BASE_URL").(string),
		Port:    viper.Get("DB_PORT").(int),
		DBName:  viper.Get("DB_NAME").(string),
		SSLMode: viper.Get("SSL_MODE").(string),
	}

	// "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	var dsn string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Jakarta",
		db.Addr, db.User, db.Passwd, db.DBName, db.Port, db.SSLMode)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to Database")
	}
}
