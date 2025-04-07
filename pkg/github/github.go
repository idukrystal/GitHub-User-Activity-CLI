package github

import "fmt"

type EventType string

const (
	PushEvent = "PushEvent"
	PullRequestEvent = "PullRequestEvent"
	ForkEvent = "ForkEvent"
	CreateEvent = "CreateEvent"
	
	
)

// Github retuns a data structure with a message if username not found
type InvalidResponse struct{
	Message string
}

// Github api retuns a list of events if username is found
type Event struct {
	Type string
	Repo Repo
	Payload Payload
}

func (e Event) GetNotificationMessage() (string, error) {
	switch eventType := e.Type; eventType {
	case PushEvent:
		return fmt.Sprintf("Pushed %d commit(s) to %s", len(e.Payload.Commits), e.Repo.Name), nil
	case PullRequestEvent:
		return fmt.Sprintf("Pull Request(%s) in %s", e.Payload.Action, e.Repo.Name), nil
	case ForkEvent:
		return fmt.Sprintf("Forked %s from %s", e.Repo.Name, e.Payload.Forkee.FullName), nil
	case CreateEvent:
		return fmt.Sprintf("Created a new %s %s in %s", e.Payload.RefType, e.Payload.Ref, e.Repo.Name), nil
	default:
		return eventType, fmt.Errorf("Unsupported Event Type: %s", eventType)
		
	}
}

func (event Event) String() string {
	str, _ := event.GetNotificationMessage()
	return str
}



type Repo struct {
	Name string
}

type Payload struct {
	Action string
	Commits []any
	Ref string
	RefType string `json:"ref_type"`
	Forkee Forkee
}

type Forkee struct {
	FullName string `json:"full_name"`
}
