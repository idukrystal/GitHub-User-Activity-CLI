package main

import (
	"fmt"
	"net/http"
	"os"
)

const NoUsernameError = "Usename Not Provided"

func main() {
	args := os.Args
	defer func() {
		if r := recover(); r != nil {
			// args[0] is the name of this program
			fmt.Println(args[0], ": ", r)
		}
	}()

	// if only one arg is provided: no username was entered
	if len(args) < 2 {
		panic(NoUsernameError)
	}

	// args[1] is the provided username
	userActivityUrl := fmt.Sprintf("https://api.github.com/users/%s/events", args[1])
	
	res , err := http.Get(userActivityUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
