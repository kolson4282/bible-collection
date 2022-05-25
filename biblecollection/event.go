package biblecollection

type Event struct {
	ID         int
	ActorID    string
	Action     string
	ReceiverID string
	Location   string
	BookID     string
	Chapter    int
	Verse      int
}

type EventCollection interface {
	GetAllEvents() []*Event
}
