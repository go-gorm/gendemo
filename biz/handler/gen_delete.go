package handler

import (
	"context"

	"github.com/go-gorm/gendemo/biz/dal/query"
	"github.com/go-gorm/gendemo/logs"
)

// SimpleDelete 简单的删除
func SimpleDelete(ctx context.Context) {
	vd := query.User
	//指定条件
	result, err := vd.WithContext(ctx).Where(vd.Name.Eq("1")).Delete()
	if err != nil {
		logs.CtxError(ctx, "[SimpleDelete] Delete err=%s", err.Error())
		return
	}
	if result.RowsAffected == 0 {
		//no deleted record
	}

	//如果你的结构里有软删除字段，默认逻辑都是走软删除
	//具体的可以参考 https://gorm.io/docs/delete.html#Find-soft-deleted-records
	//软删除的情况下，也是可以实现物理删除的,加上 Unscoped()
	result, err = vd.WithContext(ctx).Where(vd.Name.Eq("1")).Unscoped().Delete()
	if err != nil {
		logs.CtxError(ctx, "[SimpleDelete] Delete err=%s", err.Error())
		return
	}
}
