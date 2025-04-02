package github

import "fmt"

type EventType string

const (
	PushEvent = "PushEvent"
	CreateEvent = "CreateEvent"
)

type Event struct {
	Type string
	Repo Repo
	Payload Payload
}

func (e Event) GetNotificationMessage() (string, error) {
	switch eventType := e.Type; eventType {
	case PushEvent:
		return fmt.Sprintf("Pushed %d commits to %s", len(e.Payload.Commits), e.Repo.Name), nil
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
	Commits []any
	Ref string
	RefType string `json:"ref_type"`
}

