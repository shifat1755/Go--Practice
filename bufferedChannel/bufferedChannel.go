package main

import (
	"fmt"
)

func main(){
	c:= make(chan string,2)
	c<-"Dog"
	c<-"Bites"
	msg:= <-c
	fmt.Println(msg)
	msg = <-c
	fmt.Println(msg)

}