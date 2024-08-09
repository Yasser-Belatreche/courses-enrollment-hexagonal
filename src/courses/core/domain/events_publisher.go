package domain

type EventsPublisher interface {
	Publish(events []CourseEvent[interface{}]) error
}
