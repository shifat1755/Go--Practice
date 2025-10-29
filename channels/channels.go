package main

import (
	"fmt"
	"time"
)
func main(){
	c :=make(chan string)
	go count("Dog",c)
	for {
		msg,open:= <-c
		if !open{
			break
		}
		fmt.Println(msg)
	}
}
func count(name string, c chan string){
	for i:=1; i<=5; i++{
		c <- name
		time.Sleep(time.Millisecond*500)
	}
	close(c)
}