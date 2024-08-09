package courses

import (
	"github.com/Yasser-Belatreche/courses-enrollment-hexagonal/src/courses/core/domain"
	"github.com/Yasser-Belatreche/courses-enrollment-hexagonal/src/courses/core/usecases"
)

type Facade struct {
	CourseRepository     domain.CourseRepository
	EnrollmentRepository domain.EnrollmentRepository
	EventsPublisher      domain.EventsPublisher
}

func (f *Facade) ScheduleCourse(command *usecases.ScheduleCourseCommand) (*usecases.ScheduleCourseCommandResponse, error) {
	handler := &usecases.ScheduleCourseCommandHandler{
		CourseRepository: f.CourseRepository,
		EventsPublisher:  f.EventsPublisher,
	}

	return handler.Handle(command)
}

func (f *Facade) CancelCourse(command *usecases.CancelCourseCommand) (*usecases.CancelCourseCommandResponse, error) {
	handler := &usecases.CancelCourseCommandHandler{}

	return handler.Handle(command)
}

func (f *Facade) Enroll(command *usecases.EnrollCommand) (*usecases.EnrollCommandResponse, error) {
	handler := &usecases.EnrollCommandHandler{}

	return handler.Handle(command)
}

func (f *Facade) CancelEnrollment(command *usecases.CancelEnrollmentCommand) (*usecases.CancelEnrollmentCommandResponse, error) {
	handler := &usecases.CancelEnrollmentCommandHandler{}

	return handler.Handle(command)
}

func (f *Facade) GetCourse(query *usecases.GetCourseQuery) (*usecases.GetCourseQueryResponse, error) {
	handler := &usecases.GetCourseQueryHandler{
		CourseRepository:     f.CourseRepository,
		EnrollmentRepository: f.EnrollmentRepository,
	}

	return handler.Handle(query)
}
