package main

import (
	"errors"
	"sync"
)

/**
模拟队列
 */
type element interface {}
type Queue struct {
	Elements []element
}
var lock sync.Mutex

type MyQueue interface {
	Get() (element,error)
	Add(e element)
	IsEmpty() bool
}

/**
获取头部的
 */
func (q *Queue)Get() (element,error) {
	lock.Lock()
	n := len(q.Elements)
	if n == 0 {
		return nil,errors.New("队列为空")
	}
	defer lock.Unlock()
	e := q.Elements[0]
	q.Elements = q.Elements[1:]
	return e,nil
}
func (q *Queue)Add(e element){
	lock.Lock()
	q.Elements = append(q.Elements,e)
	defer lock.Unlock()
}
func (q *Queue)IsEmpty() bool {
	lock.Lock()
	n := len(q.Elements)
	if n == 0 {
		return true
	}
	defer lock.Unlock()
	return false
}
