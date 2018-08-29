package main

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_"github.com/go-sql-driver/mysql"
	"testing"
	"strconv"
)

func Test(t *testing.T)  {
	sql:=fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8","debian-sys-maint", "aNNcNHBS2HHyIQMD","127.0.0.1","3306","howie")
	db, err := xorm.NewEngine("mysql", sql)
	fmt.Println(err)
	fmt.Println(db.Ping())
	db.ShowSQL(false)
	for i:=0;i<=1000000;i++{
		db.Exec(fmt.Sprintf("INSERT INTO test_user (account, password) VALUES (%s,%s)",strconv.Itoa(i),strconv.Itoa(i)))
	}
}
