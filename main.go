package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type eventType string

type Event struct {
	Type eventType
	Repo Repo
	Payload Payload
}

type Repo struct {
	Name string
}

type Payload struct {
	Commits []any
}

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

	// the response body must be closed after the function completes
	defer res.Body.Close()

	var events []Event

	if err := json.NewDecoder(res.Body).Decode(&events); err != nil {
		panic(err)
	}

	for _, event := range events {
		fmt.Println(event.Type)
	}
}
