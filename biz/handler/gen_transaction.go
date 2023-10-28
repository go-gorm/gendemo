package handler

import (
	"context"

	"github.com/go-gorm/gendemo/biz/dal/model"
	"github.com/go-gorm/gendemo/biz/dal/query"
	"github.com/go-gorm/gendemo/logs"
)

// AutoTransAction 自动事务
func AutoTransAction(ctx context.Context) {
	//通过Transaction，开事务，只要返回err就会自动回滚，没有err就会自动提交
	err := query.Q.Transaction(func(tx *query.Query) error {
		//事务里要用事务的query操作（tx）
		err := tx.Role.WithContext(ctx).Create(&model.Role{Name: "test"})
		if err != nil { //有err返回
			return err
		}
		err = tx.User.WithContext(ctx).Create(&model.User{Name: "1"})
		if err != nil { //有err返回自动回滚
			return err
		}
		return err
	})
	if err != nil {
		logs.CtxError(ctx, "[AutoTransAction] Transaction err=%s", err.Error())
	}

	//事务也支持嵌套
	err = query.Q.Transaction(func(tx *query.Query) error {
		//事务里要用事务的query操作（tx）
		cerr := tx.Role.WithContext(ctx).Create(&model.Role{Name: "test"})
		if cerr != nil { //有err返回
			return cerr
		}
		tx.Transaction(func(tx2 *query.Query) error {
			uerr := tx2.User.WithContext(ctx).Create(&model.User{Name: "1"})
			if uerr != nil { //有err返回自动回滚
				return uerr
			}
			return tx2.User.WithContext(ctx).Create(&model.User{Name: "2"})
		})

		return err
	})
	if err != nil {
		logs.CtxError(ctx, "[AutoTransAction] Transaction err=%s", err.Error())
	}
}

// ManualTransAction 手动事务
func ManualTransAction(ctx context.Context) {
	//开启事务
	tx := query.Q.Begin()
	//记得判断err
	if tx.Error != nil {
		//事务开失败了
		return
	}
	var err error
	defer func() {
		if recover() != nil || err != nil {
			//回滚事务
			_ = tx.Rollback()
		}
	}()

	err = tx.Role.WithContext(ctx).Create(&model.Role{Name: "test"})
	if err != nil {
		return
	}
	//提交事务
	err = tx.Commit()

}
