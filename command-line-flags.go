package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")
	numPtr := flag.Int("numb", 42, "a int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	flag.Parse()

	fmt.Println("word: ", *wordPtr)
	fmt.Println("number:", *numPtr)
	fmt.Println("boolean: ", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail", flag.Args())
	http.HandleFunc()
}
