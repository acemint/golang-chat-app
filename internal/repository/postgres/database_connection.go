package repository

import (
	"fmt"
	"strconv"

	"chat-app/internal/configuration"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var MemberRepository *MemberRepositoryStruct

func InitializeConnection() error {
	driver := setupDriver()
	created_db, err := gorm.Open(driver, &gorm.Config{})
	if err != nil {
		return err
	}
	DB = created_db

	initializeRepositories()
	return nil
}

func setupDriver() gorm.Dialector {
	dbConf := configuration.AppConfiguration.Database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbConf.Host,
		dbConf.Username,
		dbConf.Password,
		dbConf.DatabaseName,
		strconv.Itoa(int(dbConf.Port)),
		dbConf.SSLMode,
		dbConf.TimeZone)
	return postgres.New(postgres.Config{DSN: dsn, PreferSimpleProtocol: true})
}

func initializeRepositories() {
	MemberRepository = &MemberRepositoryStruct{
		db: DB,
	}
}
