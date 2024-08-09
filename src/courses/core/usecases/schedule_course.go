package usecases

import "github.com/Yasser-Belatreche/courses-enrollment-hexagonal/src/courses/core/domain"

type ScheduleCourseCommand struct {
	Name    string
	MaxSize int
	MinSize int
}

type ScheduleCourseCommandResponse struct {
	Id string
}

type ScheduleCourseCommandHandler struct {
	CourseRepository domain.CourseRepository
	EventsPublisher  domain.EventsPublisher
}

func (h *ScheduleCourseCommandHandler) Handle(command *ScheduleCourseCommand) (*ScheduleCourseCommandResponse, error) {
	course, err := domain.ScheduleCourse(command.Name, command.MinSize, command.MaxSize)
	if err != nil {
		return nil, err
	}

	err = h.CourseRepository.Create(course)
	if err != nil {
		return nil, err
	}

	err = h.EventsPublisher.Publish(course.PullEvents())
	if err != nil {
		return nil, err
	}

	return &ScheduleCourseCommandResponse{Id: course.State().Id}, nil
}
