package main

import "fmt"

type close func()

func invokeClose(c close){ //c es una funcion
	c()
}

func main(){
	f := func(){
		fmt.Println("hello world")
	}
	invokeClose(f)
}