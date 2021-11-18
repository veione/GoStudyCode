package main

import (
	"fmt"
	"sync"
	"time"
)

type sp interface {
	Out(key string, val interface{})                  //存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
	Rd(key string, timeout time.Duration) interface{} //读取一个key，如果key不存在阻塞，等待key存在或者超时
}

type Entry struct {
	data    interface{}
	ch      chan struct{}
	isExist bool
}

type Map struct {
	m   map[string]*Entry
	rmx sync.RWMutex
}

func (mp Map) Out(key string, val interface{}) {
	mp.rmx.Lock()
	defer mp.rmx.Unlock()

	if d, ok := mp.m[key]; ok {
		d.isExist = true
		d.data = val
		close(d.ch)
	} else {
		temp := &Entry{
			data:    val,
			ch:      make(chan struct{}),
			isExist: true,
		}
		mp.m[key] = temp
		close(temp.ch)
	}
}

func (mp Map) Rd(key string, timeout time.Duration) interface{} {
	mp.rmx.RLocker()
	defer mp.rmx.RUnlock()

	if d, ok := mp.m[key]; ok {
		if d.isExist {
			return d.data
		} else {
			select {
			case <-d.ch:
				{
					return mp.m[key].data
				}
			case <- time.After(timeout):
				{
					fmt.Println("超时")
					return nil
				}
			}
		}
	} else {
		temp := &Entry{
			ch:      make(chan struct{}),
			isExist: false,
		}
		mp.m[key] = temp
		select {
		case <-temp.ch:
			{
				return mp.m[key].data
			}
		case <- time.After(timeout):
			{
				fmt.Println("超时")
				return nil
			}
		}
	}

}
