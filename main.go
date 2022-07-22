package main

import (
	"context"

	"github.com/go-gorm/gendemo/biz/test_model"
	"github.com/go-gorm/gendemo/mysql"
)

func main() {
	mysql.Init()
	db := mysql.DB(context.Background())
	db.AutoMigrate(test_model.JustUser{}, test_model.Person{}, test_model.Addr{})
}
