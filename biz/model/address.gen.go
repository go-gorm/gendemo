// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameAddr = "address"

// Addr mapped from table <address>
type Addr struct {
	ID     int64   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id" newTag:"tag info"`
	Street *string `gorm:"column:street" json:"street"`
	UserID *int64  `gorm:"column:user_id" json:"user_id"`
}

// TableName Addr's table name
func (*Addr) TableName() string {
	return TableNameAddr
}
