package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	arg := flag.Arg(0)
	fmt.Printf("Hello %s\n", arg)

	a := "hello"
	b := &a
	fmt.Println(*b)
	*b = "こんにちは"
	fmt.Println(a)
}
