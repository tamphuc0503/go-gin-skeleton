package common

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func CreateDatabase(driver string, database string, host string, user string, password string, port int) (*gorm.DB, error) {
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	var dialect gorm.Dialector = CreateDialect(driver, database, host, user, password, port)

	db, err := gorm.Open(dialect, gormConfig)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(20))

	return db, nil
}

func CreateDialect(driver string, database string, host string, user string, password string, port int) (gorm.Dialector){
	var dialect gorm.Dialector = nil
	if driver == "postgres" {
		dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
			host,
			port,
			user,
			database,
			password,
			"disable",
		)

		dialect = postgres.Open(dsn)
	} else if driver == "mysql" {
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			user,
			password,
			host,
			port,
			database,
		)

		dialect = mysql.Open(dsn)
	} else if driver == "sqlite" {
		dialect = sqlite.Open(database)
	}
	return dialect
}