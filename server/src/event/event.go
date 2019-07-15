package event

//type ID string

type Event struct {
	Name         string `firestore:"name"`
	PeopleJoined int    `firestore:"peopleJoined"`
}

// New creates a new, unrouted event.
func New(name string) *Event {
	return &Event{
		Name:         name,
		PeopleJoined: 0,
	}
}
