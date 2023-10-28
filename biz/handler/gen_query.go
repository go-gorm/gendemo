package handler

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-gorm/gendemo/biz/dal/model"
	"github.com/go-gorm/gendemo/biz/dal/query"
	"github.com/go-gorm/gendemo/logs"
	"gorm.io/datatypes"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

// SingleValueQuery 单值查询
func SingleValueQuery(ctx context.Context) {
	u := query.User

	// 查用户1的第一条
	user, err := u.WithContext(ctx).Where(u.ID.Eq(1)).First()
	// SELECT * FROM user  WHERE `user`.`id` = 1 ORDER BY id LIMIT 1;
	if err != nil {
		logs.CtxError(ctx, "[SingleValueQuery] First err=%s", err.Error())
	}
	fmt.Println(user)

	// 查一条，走数据库的默认排序
	user, err = u.WithContext(ctx).Take()
	// SELECT * FROM users LIMIT 1;
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logs.CtxError(ctx, "[SingleValueQuery] Take err=%s", err.Error())
	}
	fmt.Println(user)

	// 查用户1的最后一条
	user, err = u.WithContext(ctx).Where(u.ID.Eq(1)).Last()
	// SELECT * FROM user  WHERE `user`.`id` = 1 ORDER BY id DESC LIMIT 1;
	if err != nil {
		logs.CtxError(ctx, "[SingleValueQuery] Last err=%s", err.Error())
	}
	fmt.Println(user)

	// 走主库查询
	user, err = u.WithContext(ctx).WriteDB().Last()

	// check error ErrRecordNotFound
	errors.Is(err, gorm.ErrRecordNotFound)

	//如果不希望返回err，可以用limit+find/scan
	users, err := u.WithContext(ctx).Where(u.ID.Eq(1)).Limit(1).Find()
	if len(users) > 0 {
		fmt.Println(users[0])
	}

	var User *model.User
	err = u.WithContext(ctx).Where(u.ID.Eq(1)).Limit(1).Scan(&User)
	fmt.Println(User)
}

// SimpleQuery 简单查询
func SimpleQuery(ctx context.Context) {
	u := query.User

	//单条件,用户1的数据
	users, err := u.WithContext(ctx).Where(u.ID.Eq(1)).Limit(10).Find()
	if err != nil {
		logs.CtxError(ctx, "[SimpleQuery] Find err=%s", err.Error())
	}
	fmt.Println(users)

	//统计
	total, err := u.WithContext(ctx).Where(u.ID.Eq(1)).Count()
	if err != nil {
		logs.CtxError(ctx, "[SimpleQuery] Count err=%s", err.Error())
	}
	fmt.Println(total)

	//分页查询,排序
	users, total, err = u.WithContext(ctx).Where(u.ID.Eq(1)).Order(u.ID.Desc(), u.Phone).FindByPage(0, 10)
	if err != nil {
		logs.CtxError(ctx, "[SimpleQuery] Count err=%s", err.Error())
	}
	fmt.Println(total)
	fmt.Println(users)

	//多条件AND查询 用户1
	users, total, err = u.WithContext(ctx).Where(u.Name.Eq("1")).FindByPage(0, 10)

	//OR条件查询
	// （`user`.`id` = 1 AND `user`.`name` = 1) OR (`user`.`id` = 2)
	users, total, err = u.WithContext(ctx).Where(u.ID.Eq(1), u.Name.Eq("1")).Or(u.ID.Eq(2)).FindByPage(0, 10)

	// 元组查询，多个字段的条件组合
	//查询用户1类型 以及用户2类型2的
	users, err = u.WithContext(ctx).Where(u.Columns(u.Name, u.Phone).In(field.Values([][]interface{}{{"1", "1"}, {"2", "2"}}))).Find()
	// SELECT * FROM `user` WHERE (`user.name`, `user.phone`) IN ((1,1),(2,2);

	//JSON查询
	users, err = u.WithContext(ctx).Where(gen.Cond(datatypes.JSONQuery("extra").HasKey("xx_id"))...).Find()
	// SELECT * FROM `user` WHERE JSON_EXTRACT(`attributes`,'$.user_id') IS NOT NULL;

	//选择查询的字段
	users, err = u.WithContext(ctx).Select(u.ID, u.Name).Where(u.ID.Eq(1)).Find()
	//去重
	users, err = u.WithContext(ctx).Distinct(u.ID, u.Phone).Order(u.ID, u.Name.Desc()).Find()

	//分组
	var Tuser []struct {
		Name  string
		Total int
	}
	err = u.WithContext(ctx).Select(u.ID, u.ID.Count().As("total")).Group(u.ID).Scan(&Tuser)
	//加上Having
	err = u.WithContext(ctx).Select(u.ID, u.ID.Count().As("total")).Group(u.ID).Having(u.Name.Eq("10")).Scan(&Tuser)

}

// ComplexQuery 复杂查询
func ComplexQuery(ctx context.Context) {
	u := query.User
	c := query.Role
	//连表查询（示例没有业务含义）
	type Result struct {
		RoleName string
		Phone    string
		ID       int64
	}

	var result Result
	//和其他表连接
	err := u.WithContext(ctx).Select(u.ID, u.Phone, c.Name.As("role_name")).LeftJoin(c, c.ID.EqCol(u.RoleID)).Scan(&result)
	// SELECT `user`.`phone`,`role`.`name` as role_name FROM `user` left join `role` on `role`.`id` = `user`.`role_id`
	if err != nil {
		logs.CtxError(ctx, "[ComplexQuery] LeftJoin Scan err=%s", err.Error())
	}
	fmt.Println(result)

	// 同一个表做连接
	var result2 Result
	u2 := u.As("u2")
	err = u.WithContext(ctx).Select(u.ID, u2.Phone).LeftJoin(u2, u2.ID.EqCol(u.ID)).Scan(&result2)
	// SELECT users.id, u2.phone FROM `user` left join `user` u2 on u2.id = user.id

	//子查询的连接
	var result3 Result
	c2 := c.As("c2")
	err = u.WithContext(ctx).Select(u.ID, c2.Name).LeftJoin(c.WithContext(ctx).Select(c.ID, c.Name).Where(c.ID.Gt(100)).As("c2"), c2.ID.EqCol(u.RoleID)).Scan(&result3)
	// SELECT `user`.`id`,c2.`name` FROM `user` left join (select id,name from role  where id > 100) as c2 on c2.id = `user`.`role_id`

}
