package controllers

import (
	"database/sql"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
	r "github.com/revel/revel"
)

type GormController struct {
	*r.Controller
	Txn *gorm.DB
}

//Gdb export db connection
var Gdb *gorm.DB

//InitDB init mysql connection
func InitDB() {
	var err error
	// open db
	Gdb, err = gorm.Open("mysql", "root:Qwerty123456.@/goshop?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

}

// Begin transaction
func (c *GormController) Begin() r.Result {
	txn := Gdb.Begin()
	if txn.Error != nil {
		panic(txn.Error)
	}
	c.Txn = txn
	return nil
}

// Commit transaction
func (c *GormController) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}
	c.Txn.Commit()
	if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

// Rollback transaction
func (c *GormController) Rollback() r.Result {
	if c.Txn == nil {
		return nil
	}
	c.Txn.Rollback()
	if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
