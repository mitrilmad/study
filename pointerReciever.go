package main

import "fmt"

type Man struct {
	money int
}

func (man *Man) test1(money int) {
	man.money = money
}

func (man Man) test2(money int) {
	man.money = money
}

func main() {
	man1 := Man{10}
	man2 := Man{20}

	man1.test1(11)
	man2.test2(22)

	fmt.Println(man1)
	fmt.Println(man2)
}
