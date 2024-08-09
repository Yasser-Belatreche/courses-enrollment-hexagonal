package usecases

import "github.com/Yasser-Belatreche/courses-enrollment-hexagonal/src/courses/core/domain"

type CancelEnrollmentCommand struct {
	CourseId     string
	EnrollmentId string
}

type CancelEnrollmentCommandResponse struct {
	CourseId     string
	EnrollmentId string
}

type CancelEnrollmentCommandHandler struct {
	CourseRepository     domain.CourseRepository
	EnrollmentRepository domain.EnrollmentRepository
	EventsPublisher      domain.EventsPublisher
}

func (h *CancelEnrollmentCommandHandler) Handle(command *CancelEnrollmentCommand) (*CancelEnrollmentCommandResponse, error) {
	course, err := h.CourseRepository.FindById(command.CourseId)
	if err != nil {
		return nil, err
	}

	enrollment, err := h.EnrollmentRepository.FindById(command.EnrollmentId)
	if err != nil {
		return nil, err
	}

	err = course.CancelEnrollment(enrollment, func(enrollment *domain.Enrollment) error {
		return h.EnrollmentRepository.Delete(enrollment.Id)
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

	return &CancelEnrollmentCommandResponse{
		CourseId:     course.State().Id,
		EnrollmentId: enrollment.Id,
	}, nil
}
