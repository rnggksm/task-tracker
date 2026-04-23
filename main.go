package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: app <arg1> <arg2> ...")
		return
	}

	fmt.Println("All args:", os.Args[1:])
}
