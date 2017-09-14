package main

import (
	"fmt"
	"os"
	"flag"
)

func main() {
	a := os.Args
	fmt.Println(a)

	test := flag.String("테스트", "어처구니", "이딴건 왜 제공")
	fmt.Printf("commandline test %s\n", *test)
}
