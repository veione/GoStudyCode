package main

import "sync"

var ins *singleton
var once sync.Once
func GetInstance() *singleton {

	once.Do(func() {
		instance = &singleton{}
	})
	return ins
}
