package main

import (
	"context"

	"github.com/go-gorm/gendemo/biz/method"
	"github.com/go-gorm/gendemo/biz/model"
	"github.com/go-gorm/gendemo/mysql"
	"gorm.io/gen"
)

// unexported struct will be ignored.
type test struct {
	id  int
	Xxx string
	Ttt int
}

// generate code
func main() {
	//init db
	mysql.Init()
	db := mysql.DB(context.Background()).Debug()
	// specify the output directory (default: "./query")
	// ### if you want to query without context constrain, set mode gen.WithoutContext ###
	cfg := gen.Config{
		OutPath: "../../biz/dal",
		//Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		Mode:          gen.WithDefaultQuery,
		FieldNullable: true,
	}

	g := gen.NewGenerator(cfg)

	// reuse the database connection in Project or create a connection here
	// if you want to use GenerateModel/GenerateModelAs, UseDB is necessary or it will panic
	g.UseDB(db)

	// apply basic crud api on structs or table models which is specified by table name with function
	// GenerateModel/GenerateModelAs. And generator will generate table models' code when calling Execute.
	// g.ApplyBasic(model.User{}, g.GenerateModel("company"), g.GenerateModelAs("people", "Person", gen.FieldIgnore("address")))
	g.ApplyBasic(model.Customer{}, model.CreditCard{}, model.Bank{}) // Associations
	g.ApplyBasic(g.GenerateModelAs("user", "JustUser"))
	g.ApplyBasic(model.Passport{}, g.GenerateModel("people"),
		g.GenerateModelAs("address", "Addr",
			gen.FieldIgnore("deleted_at"),
			gen.FieldNewTag("id", `newTag:"tag info"`),
		),
	)

	// apply diy interfaces on structs or table models
	g.ApplyInterface(func(method.Method) {}, &model.Company{}, model.User{}, test{}) // struct test will be ignored
	g.ApplyInterface(func(method method.UserMethod) {}, model.User{})

	// execute the action of code generation
	g.Execute()
}
