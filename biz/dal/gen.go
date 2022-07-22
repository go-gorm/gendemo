// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

var (
	Q          = new(Query)
	Addr       *addr
	Bank       *bank
	Company    *company
	CreditCard *creditCard
	Customer   *customer
	JustUser   *justUser
	Passport   *passport
	Person     *person
	User       *user
)

func SetDefault(db *gorm.DB) {
	*Q = *Use(db)
	Addr = &Q.Addr
	Bank = &Q.Bank
	Company = &Q.Company
	CreditCard = &Q.CreditCard
	Customer = &Q.Customer
	JustUser = &Q.JustUser
	Passport = &Q.Passport
	Person = &Q.Person
	User = &Q.User
}

func Use(db *gorm.DB) *Query {
	return &Query{
		db:         db,
		Addr:       newAddr(db),
		Bank:       newBank(db),
		Company:    newCompany(db),
		CreditCard: newCreditCard(db),
		Customer:   newCustomer(db),
		JustUser:   newJustUser(db),
		Passport:   newPassport(db),
		Person:     newPerson(db),
		User:       newUser(db),
	}
}

type Query struct {
	db *gorm.DB

	Addr       addr
	Bank       bank
	Company    company
	CreditCard creditCard
	Customer   customer
	JustUser   justUser
	Passport   passport
	Person     person
	User       user
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:         db,
		Addr:       q.Addr.clone(db),
		Bank:       q.Bank.clone(db),
		Company:    q.Company.clone(db),
		CreditCard: q.CreditCard.clone(db),
		Customer:   q.Customer.clone(db),
		JustUser:   q.JustUser.clone(db),
		Passport:   q.Passport.clone(db),
		Person:     q.Person.clone(db),
		User:       q.User.clone(db),
	}
}

type queryCtx struct {
	Addr       *addrDo
	Bank       *bankDo
	Company    *companyDo
	CreditCard *creditCardDo
	Customer   *customerDo
	JustUser   *justUserDo
	Passport   *passportDo
	Person     *personDo
	User       *userDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Addr:       q.Addr.WithContext(ctx),
		Bank:       q.Bank.WithContext(ctx),
		Company:    q.Company.WithContext(ctx),
		CreditCard: q.CreditCard.WithContext(ctx),
		Customer:   q.Customer.WithContext(ctx),
		JustUser:   q.JustUser.WithContext(ctx),
		Passport:   q.Passport.WithContext(ctx),
		Person:     q.Person.WithContext(ctx),
		User:       q.User.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
