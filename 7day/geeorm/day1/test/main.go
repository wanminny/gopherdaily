package main

import (
	"day1/geeorm"

	"day1/log"
	_ "github.com/mattn/go-sqlite3"
	"k8s.io/klog/v2"
)

func main() {
	log.SetLoggerLevel(0)
	eng, err := geeorm.NewEngine("sqlite3", "myorm.db")
	if err != nil {
		log.Errorln(err)
		return
	}
	defer func() {
		eng.Close()
	}()

	s := eng.NewSession()

	//s.Raw("drop table if exists user").Exec()

	s.Raw("create table user(name string, age int)").Exec()

	relt, err := s.Raw("insert into user(name, age) values (?, ?)", "tomi", 20).Exec()
	if err != nil {
		log.Errorf("insert failed: %v", err)
	}
	klog.Infoln(relt.RowsAffected())

	//rows, err := s.Raw("select * from user").QueryRows()
	//if err != nil {
	//	log.Errorf("query failed: %v", err)
	//}
	//
	//for rows.Next() {
	//	var name string
	//	var age int
	//	rows.Scan(&name, &age)
	//	log.Infof("name: %s, age: %d", name, age)
	//	//fmt.Printf("name: %s, age: %d\n", name, age)}
	//
	//}
}
