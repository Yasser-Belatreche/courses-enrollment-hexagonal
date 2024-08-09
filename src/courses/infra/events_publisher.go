package infra

import (
	"github.com/Yasser-Belatreche/courses-enrollment-hexagonal/src/courses/core/domain"
)

func NewFakeEventsPublisher() domain.EventsPublisher {
	return &fakeEventsPublisher{}
}

type fakeEventsPublisher struct{}

func (f *fakeEventsPublisher) Publish(events []domain.CourseEvent[interface{}]) error {
	return nil
}
