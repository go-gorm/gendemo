// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/go-gorm/gendemo/biz/dal/model"
)

func newUser2(db *gorm.DB, opts ...gen.DOOption) user2 {
	_user2 := user2{}

	_user2.user2Do.UseDB(db, opts...)
	_user2.user2Do.UseModel(&model.User2{})

	tableName := _user2.user2Do.TableName()
	_user2.ALL = field.NewAsterisk(tableName)
	_user2.ID = field.NewInt64(tableName, "id")
	_user2.Name = field.NewString(tableName, "name")
	_user2.Extra = field.NewString(tableName, "extra")

	_user2.fillFieldMap()

	return _user2
}

type user2 struct {
	user2Do user2Do

	ALL   field.Asterisk
	ID    field.Int64
	Name  field.String
	Extra field.String

	fieldMap map[string]field.Expr
}

func (u user2) Table(newTableName string) *user2 {
	u.user2Do.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u user2) As(alias string) *user2 {
	u.user2Do.DO = *(u.user2Do.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *user2) updateTableName(table string) *user2 {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.Name = field.NewString(table, "name")
	u.Extra = field.NewString(table, "extra")

	u.fillFieldMap()

	return u
}

func (u *user2) WithContext(ctx context.Context) IUser2Do { return u.user2Do.WithContext(ctx) }

func (u user2) TableName() string { return u.user2Do.TableName() }

func (u user2) Alias() string { return u.user2Do.Alias() }

func (u user2) Columns(cols ...field.Expr) gen.Columns { return u.user2Do.Columns(cols...) }

func (u *user2) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *user2) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 3)
	u.fieldMap["id"] = u.ID
	u.fieldMap["name"] = u.Name
	u.fieldMap["extra"] = u.Extra
}

func (u user2) clone(db *gorm.DB) user2 {
	u.user2Do.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u user2) replaceDB(db *gorm.DB) user2 {
	u.user2Do.ReplaceDB(db)
	return u
}

type user2Do struct{ gen.DO }

type IUser2Do interface {
	gen.SubQuery
	Debug() IUser2Do
	WithContext(ctx context.Context) IUser2Do
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUser2Do
	WriteDB() IUser2Do
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUser2Do
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUser2Do
	Not(conds ...gen.Condition) IUser2Do
	Or(conds ...gen.Condition) IUser2Do
	Select(conds ...field.Expr) IUser2Do
	Where(conds ...gen.Condition) IUser2Do
	Order(conds ...field.Expr) IUser2Do
	Distinct(cols ...field.Expr) IUser2Do
	Omit(cols ...field.Expr) IUser2Do
	Join(table schema.Tabler, on ...field.Expr) IUser2Do
	LeftJoin(table schema.Tabler, on ...field.Expr) IUser2Do
	RightJoin(table schema.Tabler, on ...field.Expr) IUser2Do
	Group(cols ...field.Expr) IUser2Do
	Having(conds ...gen.Condition) IUser2Do
	Limit(limit int) IUser2Do
	Offset(offset int) IUser2Do
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUser2Do
	Unscoped() IUser2Do
	Create(values ...*model.User2) error
	CreateInBatches(values []*model.User2, batchSize int) error
	Save(values ...*model.User2) error
	First() (*model.User2, error)
	Take() (*model.User2, error)
	Last() (*model.User2, error)
	Find() ([]*model.User2, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.User2, err error)
	FindInBatches(result *[]*model.User2, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.User2) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUser2Do
	Assign(attrs ...field.AssignExpr) IUser2Do
	Joins(fields ...field.RelationField) IUser2Do
	Preload(fields ...field.RelationField) IUser2Do
	FirstOrInit() (*model.User2, error)
	FirstOrCreate() (*model.User2, error)
	FindByPage(offset int, limit int) (result []*model.User2, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUser2Do
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u user2Do) Debug() IUser2Do {
	return u.withDO(u.DO.Debug())
}

func (u user2Do) WithContext(ctx context.Context) IUser2Do {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u user2Do) ReadDB() IUser2Do {
	return u.Clauses(dbresolver.Read)
}

func (u user2Do) WriteDB() IUser2Do {
	return u.Clauses(dbresolver.Write)
}

func (u user2Do) Session(config *gorm.Session) IUser2Do {
	return u.withDO(u.DO.Session(config))
}

func (u user2Do) Clauses(conds ...clause.Expression) IUser2Do {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u user2Do) Returning(value interface{}, columns ...string) IUser2Do {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u user2Do) Not(conds ...gen.Condition) IUser2Do {
	return u.withDO(u.DO.Not(conds...))
}

func (u user2Do) Or(conds ...gen.Condition) IUser2Do {
	return u.withDO(u.DO.Or(conds...))
}

func (u user2Do) Select(conds ...field.Expr) IUser2Do {
	return u.withDO(u.DO.Select(conds...))
}

func (u user2Do) Where(conds ...gen.Condition) IUser2Do {
	return u.withDO(u.DO.Where(conds...))
}

func (u user2Do) Order(conds ...field.Expr) IUser2Do {
	return u.withDO(u.DO.Order(conds...))
}

func (u user2Do) Distinct(cols ...field.Expr) IUser2Do {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u user2Do) Omit(cols ...field.Expr) IUser2Do {
	return u.withDO(u.DO.Omit(cols...))
}

func (u user2Do) Join(table schema.Tabler, on ...field.Expr) IUser2Do {
	return u.withDO(u.DO.Join(table, on...))
}

func (u user2Do) LeftJoin(table schema.Tabler, on ...field.Expr) IUser2Do {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u user2Do) RightJoin(table schema.Tabler, on ...field.Expr) IUser2Do {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u user2Do) Group(cols ...field.Expr) IUser2Do {
	return u.withDO(u.DO.Group(cols...))
}

func (u user2Do) Having(conds ...gen.Condition) IUser2Do {
	return u.withDO(u.DO.Having(conds...))
}

func (u user2Do) Limit(limit int) IUser2Do {
	return u.withDO(u.DO.Limit(limit))
}

func (u user2Do) Offset(offset int) IUser2Do {
	return u.withDO(u.DO.Offset(offset))
}

func (u user2Do) Scopes(funcs ...func(gen.Dao) gen.Dao) IUser2Do {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u user2Do) Unscoped() IUser2Do {
	return u.withDO(u.DO.Unscoped())
}

func (u user2Do) Create(values ...*model.User2) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u user2Do) CreateInBatches(values []*model.User2, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u user2Do) Save(values ...*model.User2) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u user2Do) First() (*model.User2, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.User2), nil
	}
}

func (u user2Do) Take() (*model.User2, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.User2), nil
	}
}

func (u user2Do) Last() (*model.User2, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.User2), nil
	}
}

func (u user2Do) Find() ([]*model.User2, error) {
	result, err := u.DO.Find()
	return result.([]*model.User2), err
}

func (u user2Do) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.User2, err error) {
	buf := make([]*model.User2, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u user2Do) FindInBatches(result *[]*model.User2, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u user2Do) Attrs(attrs ...field.AssignExpr) IUser2Do {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u user2Do) Assign(attrs ...field.AssignExpr) IUser2Do {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u user2Do) Joins(fields ...field.RelationField) IUser2Do {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u user2Do) Preload(fields ...field.RelationField) IUser2Do {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u user2Do) FirstOrInit() (*model.User2, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.User2), nil
	}
}

func (u user2Do) FirstOrCreate() (*model.User2, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.User2), nil
	}
}

func (u user2Do) FindByPage(offset int, limit int) (result []*model.User2, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u user2Do) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u user2Do) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u user2Do) Delete(models ...*model.User2) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *user2Do) withDO(do gen.Dao) *user2Do {
	u.DO = *do.(*gen.DO)
	return u
}