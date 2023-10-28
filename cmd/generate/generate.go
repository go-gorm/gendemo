package main

import (
	"github.com/go-gorm/gendemo/biz/dal/model"
	"gorm.io/gen"
)

// GEN Guideline: https://gorm.io/gen/index.html

// generate code
func main() {
	//init db
	//mysql.Init()
	//db := mysql.DB(context.Background())
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
