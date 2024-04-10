package main

import (
	"k8s.io/klog/v2"
	"reflect"
)

type Account struct {
	Name     string
	Password string
}

func main() {

	klog.Infoln(reflect.ValueOf(&Account{}))
	klog.Infoln(reflect.ValueOf(&Account{}).Type())
	klog.Infoln(reflect.Indirect(reflect.ValueOf(&Account{})).Type().Name())
	// 输出: {0xc000118000 0xc000118010}
}
