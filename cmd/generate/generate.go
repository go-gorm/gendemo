package main

import (
	"strings"

	"gorm.io/gen"
	"gorm.io/gen/field"
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
	g.WithOpts(gen.FieldType("deleted_at", "gorm.DeletedAt"))
	g.WithTableNameStrategy(func(tableName string) (targetTableName string) {
		if strings.HasPrefix(tableName, "_") { //忽略下划线开头的表
			return ""
		}
		return tableName
	})
	g.WithDataTypeMap(map[string]func(columnType gorm.ColumnType) (dataType string){
		"tinyint": func(columnType gorm.ColumnType) (dataType string) {
			ct, _ := columnType.ColumnType()
			if strings.HasPrefix(ct, "tinyint(1)") {
				return "int8"
			}
			return "int16"
		},
	})

	// reuse the database connection in Project or create a connection here
	// if you want to use GenerateModel/GenerateModelAs, UseDB is necessray or it will panic
	g.UseDB(db)
	// apply basic crud api on structs or table models which is specified by table name with function
	// GenerateModel/GenerateModelAs. And generator will generate table models' code when calling Excute.
	// g.ApplyBasic(model.User{}, g.GenerateModel("company"), g.GenerateModelAs("people", "Person", gen.FieldIgnore("address")))
	// apply diy interfaces on structs or table models
	// g.ApplyInterface(func(method model.Method) {}, model.User{}, g.GenerateModel("company"))
	g.ApplyBasic(
		g.GenerateModel("user", gen.FieldModify(func(f gen.Field) gen.Field {
			if f.ColumnName == "id" {
				f.GORMTag.Remove(field.TagKeyGormDefault)
			}
			return f
		}),
			gen.FieldTag("phone", func(tag field.Tag) field.Tag {
				//tag.Set("kms","xxx")
				return tag.Set("encrypt", "xxx")
			}),

			gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
				if strings.Contains(columnName, "time") {
					return columnName + ",omitempty"
				}
				return columnName
			}),
			gen.FieldJSONTag("id", "id,string,omitempty"),
			gen.FieldIgnore("create_time"),
			gen.FieldNewTag("phone", field.Tag{"encrypt": "xxx"}),
			gen.FieldNewTagWithNS("form", func(columnName string) (tagContent string) {
				return columnName
			}),
		),
		g.GenerateModel("role", gen.FieldType("created_time", "int64"),
			gen.FieldGORMTag("extra", func(tag field.GormTag) field.GormTag {
				return tag.Set("serializer", "json")
			}),
			gen.WithMethod(gen.DefaultMethodTableWithNamer),
		),
	)

	//role := g.GenerateModel("role")
	//g.GenerateModel("user", gen.FieldRelate(field.HasMany, "Roles", role,
	//	&field.RelateConfig{
	//		// RelateSlice: true,
	//		GORMTag: field.GormTag{"foreignKey": []string{"RoleRefer"}},
	//	}),
	//)

	// execute the action of code generation
	g.Execute()
}
