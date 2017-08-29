package main

import (
	"runtime"
	"fmt"
)

func main() {

	runtime.GOMAXPROCS(1)
	done := make(chan bool, 1)
	count := 4

	go func() {
		for i:= 0; i< count ; i++  {
			done <- true
			fmt.Println("========>", i)
		}
	}()

	for i :=0; i< count +1 ; i++  {
		<-done
		fmt.Println("========>")
	}
}