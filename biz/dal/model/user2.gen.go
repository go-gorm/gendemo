// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUser2 = "user2"

// User2 mapped from table <user2>
type User2 struct {
	ID    int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name  string `gorm:"column:name;not null" json:"name"`
	Extra string `gorm:"column:extra;not null" json:"extra"`
}

// TableName User2's table name
func (*User2) TableName() string {
	return TableNameUser2
}