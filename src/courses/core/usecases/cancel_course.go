package usecases

import "github.com/Yasser-Belatreche/courses-enrollment-hexagonal/src/courses/core/domain"

type CancelCourseCommand struct {
	Id string
}

type CancelCourseCommandResponse struct {
	Id string
}

type CancelCourseCommandHandler struct {
	CourseRepository domain.CourseRepository
	EventsPublisher  domain.EventsPublisher
}

func (h *CancelCourseCommandHandler) Handle(command *CancelCourseCommand) (*CancelCourseCommandResponse, error) {
	course, err := h.CourseRepository.FindById(command.Id)
	if err != nil {
		return nil, err
	}

	err = course.Cancel()
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

	return &CancelCourseCommandResponse{Id: course.State().Id}, nil
}
