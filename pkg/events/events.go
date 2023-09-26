package events

type Events struct {
	Start,
	Stop func()
}

func New(start, stop func()) (events *Events) {
	events = new(Events)
	events.Start, events.Stop = start, stop
	return
}
