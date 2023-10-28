package main

import (
	"context"
	"strings"

	"github.com/go-gorm/gendemo/mysql"
	"gorm.io/gen"
)

// GEN Guideline: https://gorm.io/gen/index.html

// generate code
func main() {
	//init db
	mysql.Init()
	db := mysql.DB(context.Background())

	// specify the output directory (default: "./query")
	// ### if you want to query without context constrain, set mode gen.WithoutContext ###
	g := gen.NewGenerator(gen.Config{
		OutPath: "../../biz/dal/query",
		Mode:    gen.WithDefaultQuery | gen.WithQueryInterface,
		/* Mode: gen.WithoutContext,*/
		//if you want the nullable field generation property to be pointer type, set FieldNullable true
		/* FieldNullable: true,*/
	})
	g.WithTableNameStrategy(func(tableName string) (targetTableName string) {
		if strings.HasPrefix(tableName, "_") { //忽略下划线开头的表
			return ""
		}
		return tableName
	})

	// reuse the database connection in Project or create a connection here
	// if you want to use GenerateModel/GenerateModelAs, UseDB is necessray or it will panic
	g.UseDB(db)

	// apply basic crud api on structs or table models which is specified by table name with function
	// GenerateModel/GenerateModelAs. And generator will generate table models' code when calling Excute.
	// g.ApplyBasic(model.User{}, g.GenerateModel("company"), g.GenerateModelAs("people", "Person", gen.FieldIgnore("address")))
	g.ApplyBasic(g.GenerateAllTable()...) //同步数据库所有表
	// apply diy interfaces on structs or table models
	// g.ApplyInterface(func(method model.Method) {}, model.User{}, g.GenerateModel("company"))

	// execute the action of code generation
	g.Execute()
}
