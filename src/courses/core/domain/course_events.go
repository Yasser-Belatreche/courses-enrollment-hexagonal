package domain

import (
	"github.com/Yasser-Belatreche/courses-enrollment-hexagonal/src/courses/core/domain/events"
	"github.com/Yasser-Belatreche/courses-enrollment-hexagonal/src/lib"
	"time"
)

type CourseEvent[T interface{}] struct {
	Type       string
	Payload    T
	OccurredAt time.Time
	EventId    string
}

func NewCourseBecameViableEvent(course *Course) *CourseEvent[interface{}] {
	return newCourseEvent[interface{}](events.CourseBecameViableEventType, events.CourseBecameViableEventPayload{Id: course.id})
}

func NewCourseNotViableAnymoreEvent(course *Course) *CourseEvent[interface{}] {
	return newCourseEvent[interface{}](events.CourseNotViableAnymoreEventType, events.CourseNotViableAnymoreEventPayload{Id: course.id})
}

func NewCourseScheduledEvent(course *Course) *CourseEvent[interface{}] {
	return newCourseEvent[interface{}](events.CourseScheduledEventType, events.CourseScheduledEventPayload{Id: course.id})
}

func NewCourseCancelledEvent(course *Course) *CourseEvent[interface{}] {
	return newCourseEvent[interface{}](events.CourseCancelledEventType, events.CourseCancelledEventPayload{Id: course.id})
}

func NewEnrolledInCourseEvent(course *Course, enrollment *Enrollment) *CourseEvent[interface{}] {
	return newCourseEvent[interface{}](events.EnrolledInCourseEventType, events.EnrolledInCourseEventPayload{
		CourseId:     course.id,
		EnrollmentId: enrollment.Id,
		Student:      enrollment.Student,
	})
}

func NewCourseEnrollmentCancelledEvent(course *Course, enrollment *Enrollment) *CourseEvent[interface{}] {
	return newCourseEvent[interface{}](events.CourseEnrollmentCancelledEventType, events.CourseEnrollmentCancelledEventPayload{
		CourseId:     course.id,
		EnrollmentId: enrollment.Id,
		Student:      enrollment.Student,
	})
}

func newCourseEvent[T interface{}](eventType string, payload T) *CourseEvent[T] {
	return &CourseEvent[T]{
		Type:       eventType,
		Payload:    payload,
		OccurredAt: time.Now(),
		EventId:    lib.GenerateUlid(),
	}
}
