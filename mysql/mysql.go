package mysql

import (
	"context"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

const MySQLDSN = "gorm:gorm@tcp(10.37.48.128:9910)/gorm?charset=utf8mb4&parseTime=True&loc=Local"

var (
	db   *gorm.DB
	once sync.Once
)

func Init() {
	once.Do(func() {
		// init
		DB, err := gorm.Open(mysql.Open(MySQLDSN))
		// check err
		if err != nil {
			panic(err)
		}
		db = DB
		log.Printf("init mysql for %s success", MySQLDSN)
	})
}

// WriteDB ...
func WriteDB(ctx context.Context) *gorm.DB {
	return db.Clauses(dbresolver.Write).WithContext(ctx)
}

// ReadDB ...
func ReadDB(ctx context.Context) *gorm.DB {
	return db.Clauses(dbresolver.Read).WithContext(ctx)
}

// DB Read write separation
func DB(ctx context.Context) *gorm.DB {
	return db.WithContext(ctx)
}
