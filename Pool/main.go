package main

import (
	"fmt"
)

func main() {
}

func workerPool(n int,jobCh <-chan int,retCh chan<- string)  {
	for i:=0;i<n;i++ {
		go worker(i,jobCh,retCh)
	}
}
func worker(id int,jobCh <-chan int,retCh chan<- string)  {
	//直到jobCh关闭，整个循环结束，否则一直阻塞
	for job := range jobCh {
		ret := fmt.Sprintf("worker %d processed job:%d",id,job)
		retCh <- ret
	}
}

func genJob(n int) <-chan int {
	jobCh := make(chan int,200)
	go func() {
		for i:=0;i<n ;i++  {
			jobCh <- i
		}
		close(jobCh)
	}()
	return jobCh
}






