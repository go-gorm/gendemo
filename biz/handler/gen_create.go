package handler

import (
	"context"

	"github.com/go-gorm/gendemo/biz/dal/model"
	"github.com/go-gorm/gendemo/biz/dal/query"
	"github.com/go-gorm/gendemo/logs"
	"gorm.io/gorm/clause"
)

// SimpleCreate 简单的创建
func SimpleCreate(ctx context.Context) {
	vd := query.User
	user := &model.User{Name: "1", Phone: "12xxx"}
	//简单创建
	err := vd.WithContext(ctx).Create(user)
	if err != nil {
		logs.CtxError(ctx, "[SimpleCreate] Create err=%s", err.Error())
		return
	}
}

// SelectCreate 选择字段创建
func SelectCreate(ctx context.Context) {
	vd := query.User
	user := &model.User{Name: "1", Phone: "12xxx"}
	//创建语句只会写shop_id 字段
	err := vd.WithContext(ctx).Select(vd.Name).Create(user)
	if err != nil {
		logs.CtxError(ctx, "[SelectCreate] Create err=%s", err.Error())
		return
	}
}

// IgnoreCreate 忽略字段创建
func IgnoreCreate(ctx context.Context) {
	vd := query.User
	user := &model.User{Name: "1", Phone: "12xxx"}
	//创建语句忽略shop_id 字段
	err := vd.WithContext(ctx).Omit(vd.Name).Create(user)
	if err != nil {
		logs.CtxError(ctx, "[IgnoreCreate] Create err=%s", err.Error())
		return
	}
}

// BatchCreate 批量创建
func BatchCreate(ctx context.Context) {
	vd := query.User
	users := []*model.User{
		{Name: "1", Phone: "11xxx"},
		{Name: "2", Phone: "12xxx"},
		{Name: "3", Phone: "13xxx"},
	}
	//批量创建
	err := vd.WithContext(ctx).Create(users...)
	if err != nil {
		logs.CtxError(ctx, "[BatchCreate] Create err=%s", err.Error())
		return
	}
	//批量创建,设置批量创建数量2，即一次创建两条
	err = vd.WithContext(ctx).CreateInBatches(users, 2)
	if err != nil {
		logs.CtxError(ctx, "[BatchCreate] Create err=%s", err.Error())
		return
	}
}

// Upsert 更新创建
func Upsert(ctx context.Context) {
	vd := query.User
	users := []*model.User{
		{Name: "1", Phone: "11xxx"},
		{Name: "2", Phone: "12xxx"},
		{Name: "3", Phone: "13xxx"},
	}
	//创建时冲突直接忽略
	err := vd.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(users...)
	if err != nil {
		logs.CtxError(ctx, "[Upsert] Create err=%s", err.Error())
		return
	}
	//创建时冲突，更新所有字段
	err = vd.WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Create(users...)
	if err != nil {
		logs.CtxError(ctx, "[Upsert] Create err=%s", err.Error())
		return
	}
	//创建时冲突，更新指定字段
	err = vd.WithContext(ctx).Clauses(clause.OnConflict{
		DoUpdates: clause.Assignments(map[string]interface{}{"shop_id": "group_id"}),
	}).Create(users...)
	if err != nil {
		logs.CtxError(ctx, "[Upsert] Create err=%s", err.Error())
		return
	}
}
