package main

import (
	"fmt"
	//"net/http"
	"os"
)

const NoUsernameError = "Usename Not Provided"

func main() {
	args := os.Args
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(args[0], ": ", r)
		}
	}()
	if len(args) < 2 {
		panic(NoUsernameError)
	}
}
