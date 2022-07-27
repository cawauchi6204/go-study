package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// osはファイル操作?
	flag.Parse()
	arg := flag.Arg(0)
	msg := fmt.Sprintf("Hello %s\n", arg)

	f, err := os.Create("hello.txt")
	if err != nil {
		fmt.Printf("faild to create file \n: %v", err)
	}
	defer f.Close()
	_, err = f.WriteString(msg)
	if err != nil {
		fmt.Printf("faild to create file \n: %v", err)
		return
	}
}
