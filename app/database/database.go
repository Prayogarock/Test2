package database

import (
	"fmt"
	"technopartner/app/config"

	incomesData "technopartner/features/incomes/data"
	outcomesData "technopartner/features/outcomes/data"
	usersData "technopartner/features/users/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql(cfg *config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return DB
}

func InittialMigration(db *gorm.DB) {
	db.AutoMigrate(&usersData.User{})
	db.AutoMigrate(&incomesData.Income{})
	db.AutoMigrate(&outcomesData.Outcome{})
}
