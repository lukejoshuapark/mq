package channels

type Event struct {
	Type string
	Data interface{}
}

type EventBus interface {
	Subscribe() <-chan Event
	Publish(events chan<- Event) error
	Stream() chan Event
	Close(done <-chan struct{})
}
