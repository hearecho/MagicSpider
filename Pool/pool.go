package main

/**
协程池
 */

//1. 首先要定义协程池的结构
type Pool struct {
	//初始化大小
	InitCap int
	//最大协程池大小
	MaxCap int
	Tasks *Queue
	//定时清理没有再使用的协程

}
