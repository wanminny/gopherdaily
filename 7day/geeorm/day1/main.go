package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"k8s.io/klog/v2"
)

type account struct {
	id       int    `json:"id"`
	name     string `json:"name"`
	password string `json:"password"`
}

func main() {

	db, err := sql.Open("sqlite3", "gee.db")
	if err != nil {
		klog.Infoln(err.Error())
	}

	defer func() {
		db.Close()
	}()

	_, err = db.Exec("drop table if exists account")
	if err != nil {
		klog.Infoln(err.Error())
	}
	_, err = db.Exec("create table account (id integer primary key autoincrement, name text, password text)")
	if err != nil {
		klog.Infoln(err.Error())
	}
	//klog.Infoln(rlt.RowsAffected())

	// 插入数据
	_, err = db.Exec("insert into account (name, password) values ('john', '123')")
	if err != nil {
		klog.Infoln(err.Error)
	}
	//klog.Infoln(rlt.RowsAffected())

	row, err := db.Query("select * from account")
	if err != nil {
		klog.Infoln(err.Error())
	}
	var accounts []account
	for row.Next() {
		var account account
		row.Scan(&account.id, &account.name, &account.password)
		accounts = append(accounts, account)
	}
	klog.Infoln(accounts)

}
