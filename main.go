package main
// start

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/idukrystal/GitHub-User-Activity-CLI/pkg/github"
	"io"
	"net/http"
	"os"
)

type (
	Event = github.Event
	InvalidResponse = github.InvalidResponse
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

	// the response body must be closed after the function completes
	defer res.Body.Close()

	var events []Event

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(bodyBytes, &events); err != nil {
		// The github api returns a json that contains a message key to explain what went wrong
		var typeErr *json.UnmarshalTypeError
		if errors.As(err, &typeErr) {
			var invalidResponse InvalidResponse
			if err = json.Unmarshal(bodyBytes, &invalidResponse); err == nil {
				fmt.Printf("username %s: %s\n", args[1], invalidResponse.Message)
				return
			}
		}
		panic(err)
	}

	for _, event := range events {
		fmt.Println("-", event)
	}
}
