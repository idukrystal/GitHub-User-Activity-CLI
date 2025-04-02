package github

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

