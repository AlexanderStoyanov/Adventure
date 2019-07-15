package event

//type ID string

type Event struct {
	name         string `firestore:"name"`
	peopleJoined int64  `firestore:"peopleJoined"`
}

// New creates a new, unrouted event.
func New(name string) *Event {
	return &Event{
		name:         name,
		peopleJoined: 0,
	}
}
