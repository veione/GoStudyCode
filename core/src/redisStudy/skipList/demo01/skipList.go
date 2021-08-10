package main

import "math/rand"

const (
	MaxLevel = 32
	p        = 0.25
)

type Element struct {
	Score   float64
	Value   interface{}
	Forward []*Element
}

func NewElement(score float64, value interface{}, level int) *Element {
	return &Element{
		Score:   score,
		Value:   value,
		Forward: make([]*Element, level),
	}
}

type SkipList struct {
	header *Element
	level  int
	len    int
}

func NewSkipList() *SkipList {
	return &SkipList{
		header: &Element{Forward: make([]*Element, MaxLevel)},
	}
}

func RandomLevel() int {
	level := 1
	for rand.Float32() < p && level < MaxLevel {
		level++
	}
	return level
}

func (this *SkipList) Search(score float64)(e *Element, ok bool){
	temp := this.header
	for i:= this.level-1; i>=0 ;i-- {
		for temp.Forward[i] != nil && temp.Forward[i].Score < score{
			temp = temp.Forward[i]
		}
	}
	temp = temp.Forward[0]
	if temp != nil && temp.Score == score{
		return temp, true
	}
	return  nil, false
}

func (this *SkipList) Insert(score float64, value interface{}) *Element{
	update := make([]*Element, MaxLevel)
	temp := this.header
	for i:= this.level-1; i>=0 ;i-- {
		for temp.Forward[i] != nil && temp.Forward[i].Score < score{
			temp = temp.Forward[i]
		}
		update[i] = temp
	}
	temp = temp.Forward[0]
	if temp != nil && temp.Score == score{
		temp.Value =  value
		return temp
	}
	level := RandomLevel()
	if level > this.level{
		level = this.level + 1
		update[level] = this.header
		this.level = level
	}
	e := NewElement(score, value, level)
	for i:= 0; i< level ;i ++ {
		e.Forward[i] = update[i].Forward[i]
		update[i].Forward[i] = e
	}
	this.len ++
	return e
}


