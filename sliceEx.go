package main

import  _ "fmt"

func sender(sub []int){
	sub[0] = 1
	sub[1] = 1
	sub[2] = 1
	sub[3] = 2
	sub[4] = 2
	sub[5] = 2


}

func main(){
	sub := make([]int, 3, 6)
	//s := make(chan []int)

	sender(sub)

	//go sender(s, sub)

	//fmt.Println(<-s)
}
