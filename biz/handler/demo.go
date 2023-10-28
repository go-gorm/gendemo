package handler

import (
	"context"
	"fmt"

	"github.com/go-gorm/gendemo/biz/dal/model"
	"github.com/go-gorm/gendemo/biz/dal/query"
	"github.com/go-gorm/gendemo/logs"
)

func GenDemo(ctx context.Context) {
	user := query.User
	userMode := &model.User{Name: "gen", Extra: "test"}
	//创建
	err := user.WithContext(ctx).Create(userMode)
	if err != nil {
		logs.CtxError(ctx, "[GenDemo] Create err=%s", err.Error())
		return
	}
	//更新
	_, err = user.WithContext(ctx).Where(user.ID.Eq(userMode.ID)).UpdateSimple(user.Name.Value("gen_update"))
	if err != nil {
		logs.CtxError(ctx, "[GenDemo] UpdateSimple err=%s", err.Error())
		return
	}
	//查询
	users, total, err := user.WithContext(ctx).Where(user.IsDeleted.Eq(0), user.Name.Eq("gen")).FindByPage(0, 10)
	if err != nil {
		logs.CtxError(ctx, "[GenDemo] FindByPage err=%s", err.Error())
		return
	}
	for _, u := range users {
		fmt.Println(u.Name)
	}
	logs.CtxInfo(ctx, "[GenDemo] Find total=%d", total)
	//删除
	_, err = user.WithContext(ctx).Where(user.ID.Eq(userMode.ID)).Delete()
	if err != nil {
		logs.CtxError(ctx, "[GenDemo] Delete err=%s", err.Error())
		return
	}
}
