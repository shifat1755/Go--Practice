package main

import (
	"fmt"
	"sync"
)
func main()  {
	jobs :=make(chan int,100)
	res :=make(chan int,100)

	var wg sync.WaitGroup
	for i:=0;i<3;i++{
		wg.Add(1)
		go func ()  {
			worker(jobs,res)
			wg.Done()
		}()
	}

	for i:=0;i<20;i++{
		jobs<-i
	}
	close(jobs)
	go func (){
		wg.Wait()
		close(res)	
	}()
	for r:=range res{
		fmt.Println(r)
	}

}
func  worker(jobs<-chan int, res chan <-int)  {
	for n:=range jobs{
		res<-fib(n)
	}
}

func fib(n int)int{
	if n<=1{
		return n
	}
	return fib(n-1)+fib(n-2)
}