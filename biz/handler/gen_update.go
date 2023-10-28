package handler

import (
	"context"

	"github.com/go-gorm/gendemo/biz/dal/model"
	"github.com/go-gorm/gendemo/biz/dal/query"
	"github.com/go-gorm/gendemo/logs"
)

// SimpleUpdate 简单的更新
func SimpleUpdate(ctx context.Context) {
	vd := query.User
	//只更新单个字段
	update, err := vd.WithContext(ctx).Where(vd.ID.Eq(1)).Update(vd.Name, "1")
	if err != nil {
		logs.CtxError(ctx, "[SimpleUpdate] Update err=%s", err.Error())
		return
	}
	if update.RowsAffected <= 0 {
		logs.CtxWarn(ctx, "[SimpleUpdate] Update RowsAffected=0")
	}
	//更新有限多字段
	update, err = vd.WithContext(ctx).Where(vd.ID.Eq(1)).
		UpdateSimple(vd.Name.Value("2"), vd.Extra.Zero())
	if err != nil {
		logs.CtxError(ctx, "[SimpleUpdate] Update err=%s", err.Error())
		return
	}
	if update.RowsAffected <= 0 {
		logs.CtxWarn(ctx, "[SimpleUpdate] Update RowsAffected=0")
	}

	//通过struct更新，默认是会忽略零值，也就是这里只会更新 shop_id和group_id
	user := &model.User{Name: "1", Phone: "12xxx"}
	update, err = vd.WithContext(ctx).Where(vd.ID.Eq(1)).
		Updates(user)
	if err != nil {
		logs.CtxError(ctx, "[SimpleUpdate] Update err=%s", err.Error())
		return
	}
	if update.RowsAffected <= 0 {
		logs.CtxWarn(ctx, "[SimpleUpdate] Update RowsAffected=0")
	}

	//通过struct更新，还想更新非零值，可以加sleet
	update, err = vd.WithContext(ctx).Where(vd.ID.Eq(1)).
		Select(vd.Name, vd.Extra).Updates(user)
	//或者直接更新所有，select *
	update, err = vd.WithContext(ctx).Where(vd.ID.Eq(1)).
		Select(vd.ALL).Updates(user)
	if err != nil {
		logs.CtxError(ctx, "[SimpleUpdate] Update err=%s", err.Error())
		return
	}
	if update.RowsAffected <= 0 {
		logs.CtxWarn(ctx, "[SimpleUpdate] Update RowsAffected=0")
	}

	//当然除了上面的方式，也可以用map，会更新map里的所有字段
	update, err = vd.WithContext(ctx).Where(vd.ID.Eq(1)).
		Updates(map[string]interface{}{"name": "3", "phone": 3})
	if err != nil {
		logs.CtxError(ctx, "[SimpleUpdate] Update err=%s", err.Error())
		return
	}
	if update.RowsAffected <= 0 {
		logs.CtxWarn(ctx, "[SimpleUpdate] Update RowsAffected=0")
	}
}

// UpdateFromQuery 查询更新
//！！！这里只是举例如何根据查询更新，没有任何业务意义甚至看起来更新很不合理
func UpdateFromQuery(ctx context.Context) {
	vd := query.User
	co := query.Role
	update, err := vd.WithContext(ctx).Debug().Where(vd.ID.Eq(1)).
		Update(vd.Name, co.WithContext(ctx).Select(co.Name).Where(co.ID.Eq(3)))
	//UPDATE `user` SET `name`=(SELECT `role`.`name` FROM `role` WHERE `role`.`id` = 3)
	//WHERE `user`.`id` = 1
	if err != nil {
		logs.CtxError(ctx, "[UpdateFromQuery] Update err=%s", err.Error())
		return
	}
	if update.RowsAffected <= 0 {
		logs.CtxWarn(ctx, "[UpdateFromQuery] Update RowsAffected=0")
	}

	//多个字段更新
	vdt := vd.As("vdt")
	ua := vd.As("ua")
	update, err = ua.WithContext(ctx).Debug().
		UpdateFrom(vdt.WithContext(ctx).Select(vd.Name, vd.Phone).Where(vd.ID.Eq(1))).
		Where(ua.ID.Eq(2)).UpdateSimple(
		ua.Name.SetCol(vdt.Name), ua.Phone.SetCol(vdt.Phone))
	//UPDATE `user` AS `ua`,
	//(SELECT `user`.`name`,`user`.`phone` FROM `user` WHERE `user`.`id` = 1) AS `vdt`
	//SET `ua`.`name`=`vdt`.`name`,`ua`.`phone`=`vdt`.`phone` WHERE `ua`.`id` = 2
	if err != nil {
		logs.CtxError(ctx, "[UpdateFromQuery] Update err=%s", err.Error())
		return
	}
	if update.RowsAffected <= 0 {
		logs.CtxWarn(ctx, "[UpdateFromQuery] Update RowsAffected=0")
	}

}
