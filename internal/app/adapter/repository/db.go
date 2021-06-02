package repository

import (
	"os"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	maxIdleConn = 10
	maxOpenConn = 100
)

// Connection gets connection of mysql database
func Connection() (db *gorm.DB) {
	dsn := os.Getenv("dsn")
	if !strings.Contains(dsn, "?") {
		dsn += "?charset=utf8&parseTime=True"
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix:   "t_", // table name prefix, table for `User` would be `t_users`
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()

	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetMaxOpenConns(maxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}
