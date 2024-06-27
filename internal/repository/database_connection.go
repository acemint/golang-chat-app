package repository

import (
	"fmt"
	"strconv"

	"chat-app/internal/configuration"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDB *gorm.DB

// TODO: Create repository methods & transactional for DB Connections
func InitializePostgreSqlConnection() error {
	dbConf := configuration.AppConfiguration.Database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbConf.Host,
		dbConf.Username,
		dbConf.Password,
		dbConf.DatabaseName,
		strconv.Itoa(int(dbConf.Port)),
		dbConf.SSLMode,
		dbConf.TimeZone)

	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn, PreferSimpleProtocol: true}), &gorm.Config{})
	if err != nil {
		return err
	}
	PostgresDB = db
	return nil
}
