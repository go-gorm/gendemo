package query

import (
	"context"

	"github.com/go-gorm/gendemo/mysql"
)

func init() {
	mysql.Init()
	SetDefault(mysql.DB(context.Background()))
}
