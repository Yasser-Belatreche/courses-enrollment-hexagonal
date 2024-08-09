package usecases

import "github.com/Yasser-Belatreche/courses-enrollment-hexagonal/src/courses/core/domain"

type EnrollCommand struct {
	CourseId string
	Student  string
}

type EnrollCommandResponse struct {
	CourseId     string
	EnrollmentId string
}

type EnrollCommandHandler struct {
	CourseRepository     domain.CourseRepository
	EnrollmentRepository domain.EnrollmentRepository
	EventsPublisher      domain.EventsPublisher
}

func (h *EnrollCommandHandler) Handle(command *EnrollCommand) (*EnrollCommandResponse, error) {
	course, err := h.CourseRepository.FindById(command.CourseId)
	if err != nil {
		return nil, err
	}

	enrollment, err := course.Enroll(command.Student, func(enrollment *domain.Enrollment) error {
		return h.EnrollmentRepository.Create(enrollment)
	})
	if err != nil {
		return nil, err
	}

	err = h.CourseRepository.Update(course)
	if err != nil {
		return nil, err
	}

	err = h.EventsPublisher.Publish(course.PullEvents())
	if err != nil {
		return nil, err
	}

	return &EnrollCommandResponse{
		CourseId:     course.State().Id,
		EnrollmentId: enrollment.Id,
	}, nil
}
