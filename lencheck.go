package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	n := 2
	size := (len(a) + n - 1) / n
	fmt.Println(size)
}
