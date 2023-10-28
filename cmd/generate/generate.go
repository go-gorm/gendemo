package main

import (
	"github.com/go-gorm/gendemo/biz/dal/model"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/rawsql"
)

// GEN Guideline: https://gorm.io/gen/index.html

// generate code
func main() {
	//init db
	//mysql.Init()
	//db := mysql.DB(context.Background())
	db, _ := gorm.Open(rawsql.New(rawsql.Config{
		//SQL:      rawsql,                      //指定建表sql字符串也行
		FilePath: []string{
			//"./sql/user.sql", // 可以是sql文件,文件里可以是单个建表语句，也可以是多个
			"../sql/", // 指定一个sql文件路径也可以
		},
	}))
	// specify the output directory (default: "./query")
	// ### if you want to query without context constrain, set mode gen.WithoutContext ###
	g := gen.NewGenerator(gen.Config{
		OutPath: "../../biz/dal/query",
		Mode:    gen.WithDefaultQuery | gen.WithQueryInterface,
		/* Mode: gen.WithoutContext,*/
		//if you want the nullable field generation property to be pointer type, set FieldNullable true
		/* FieldNullable: true,*/
	})
	g.ApplyBasic(
		model.User{},
		model.Role{},
	)

	// execute the action of code generation
	g.Execute()
}
